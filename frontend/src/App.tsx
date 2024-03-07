import "./App.css";
import { Curl, Save, GetTpl } from "../wailsjs/go/main/App";
import { useRef, useState } from "react";
import { LogDebug, LogError } from "../wailsjs/runtime";
import { objectEntries } from "./utils";
import { useAlert } from "./Helper";

type Result = [string, string, string, string];

function App() {
  const curlDlg = useRef<HTMLDialogElement>(null);
  const [alertDlg, alert] = useAlert();
  const [curl, setCurl] = useState("");
  const [swagger, setSwagger] = useState("");
  const [result, setResult] = useState<Result>();
  const [summary, infs, funcs, defs] = result ?? ["", "", "", ""];

  return (
    <div id="App">
      <div className="flex">
        <div className="pr-2 w-50">
          swagger json: <br />
          <textarea
            className="content"
            value={swagger}
            onChange={(e) => setSwagger(e.target.value)}
          ></textarea>
        </div>
        <div className="w-50">
          output:
          <pre id="output" className="content">
            {summary}
          </pre>
        </div>
      </div>

      <div className="text-center mt-1">
        <button
          onClick={async () => {
            curlDlg.current?.showModal();
          }}
        >
          Curl
        </button>
        &nbsp;
        <button
          onClick={async () => {
            if (swagger) {
              setResult(undefined);
              const rs = await format(swagger);
              setResult(rs);
            } else {
              alert("swagger is null");
            }
          }}
        >
          Parse
        </button>
        &nbsp;
        <button
          onClick={async () => {
            if (infs) {
              await Save(infs, funcs, defs);
              alert("save done!");
            } else {
              alert("infs is null");
            }
          }}
        >
          Save
        </button>
        {/* &nbsp;
        <button
          onClick={async () => {
            const data = await Greet("foobar");
            setSwagger(data);
            LogDebug(data);
          }}
        >
          Test
        </button> */}
      </div>

      <dialog id="dlg-curl" ref={curlDlg} open={false}>
        curl: <br />
        <textarea
          rows={20}
          value={curl}
          onChange={(e) => setCurl(e.target.value)}
        ></textarea>
        <div className="text-center mt-1">
          <form method="dialog" className="inline">
            <button
              onClick={async () => {
                if (curl) {
                  setSwagger("");
                  const rs = await Curl(curl);
                  setSwagger(rs);
                  alert("curl done!");
                } else {
                  alert("curl is null");
                }
              }}
            >
              Get JSON
            </button>
            &nbsp;
            <button>OK</button>
          </form>
        </div>
      </dialog>
      {alertDlg}
    </div>
  );
}

export default App;

async function format(val: string): Promise<Result> {
  if (val) {
    try {
      const swg = JSON.parse(val);
      const summary = formatSummary(swg);
      const [infs, funcs] = await formatInfs(swg);
      const defs = formatDefs(swg);
      return [summary, infs, funcs, defs];
    } catch (error) {
      LogError((error as any)?.stack);
    }
  }
  return ["", "", "", ""];
}

function formatSummary(swg: Swagger.RootObject) {
  const pathsMap: { [k: string]: string[] } = {};
  for (const pathKey in swg.paths) {
    const superName = pathKey.substring(0, pathKey.indexOf("/", 2));
    const paths = (pathsMap[superName] ??= []);
    paths.push(pathKey);
  }
  let summary = "";
  for (const superName in pathsMap) {
    const paths = pathsMap[superName];
    summary += superName + ":\n";
    for (const path of paths) {
      summary += "  " + path + "\n";
    }
  }
  return summary;
}

interface PathItem {
  params: string[];
  body: string[];
  resp: string;
  name: string;
  method: string;
  pathItem: Swagger.PathItem;
  pathKey: string;
}

async function formatInfs(swg: Swagger.RootObject) {
  const getInfTpl = await GetTpl("get.inf.tpl");
  const getFunTpl = await GetTpl("get.fun.tpl");
  const postInfTpl = await GetTpl("post.inf.tpl");
  const postFunTpl = await GetTpl("post.fun.tpl");

  const pathsMap: { [k: string]: PathItem[] } = {};
  for (const pathKey in swg.paths) {
    const item = swg.paths[pathKey];
    const pathSegments = pathKey.split("/");
    if (pathSegments.length < 2) continue;
    const pathName = pathSegments[pathSegments.length - 1];
    const superName = pathSegments[1];
    console.debug(`formatInfs: ${superName}`);
    const paths = (pathsMap[superName] ??= []);
    objectEntries(item).forEach(([method, pathItem]) => {
      const params = getParams(pathItem.parameters ?? []);
      const body = getBody(pathItem.parameters ?? []);
      const resp = getType(pathItem.responses["200"].schema);
      const name = /* pathName.length < 10 ? pathItem.operationId : */ pathName;
      paths.push({ params, body, resp, name, pathItem, method, pathKey });
    });
  }

  let interfaces = "export interface Service {\n";
  let functions = "export const service = {\n";

  for (const superName in pathsMap) {
    const paths = pathsMap[superName];
    interfaces += `${superName}: {\n`;
    functions += `${superName}: {\n`;
    for (const {
      pathItem,
      resp,
      name,
      params,
      body,
      method,
      pathKey,
    } of paths) {
      const fromTpl = (tpl: string) => {
        return tpl
          .replaceAll("${pathItem.summary}", pathItem.summary)
          .replaceAll("${name}", name)
          .replaceAll("${resp}", resp)
          .replaceAll("${method}", method)
          .replaceAll("${pathKey}", pathKey)
          .replaceAll("${params}", params.join(","))
          .replaceAll("${body}", body.join(","));
      };
      {
        /* 接口 */
        const inf = fromTpl(getInfTpl);
        interfaces += inf;
        /* 函数 */
        const fun = fromTpl(getFunTpl);
        functions += fun;
      }
      if (body.length > 0) {
        /* 接口 */
        const inf = fromTpl(postInfTpl);
        interfaces += inf;
        /* 函数 */
        const fun = fromTpl(postFunTpl);
        functions += fun;
      }
    }
    interfaces += "\n}\n";
    functions += "\n},\n";
  }

  interfaces += "}";
  functions += "}";
  return [interfaces, functions];
}

interface Inf {
  name: string;
  props: Prop[];
}
interface Prop {
  propName: string;
  propType: string;
}
function formatDefs(swg: Swagger.RootObject) {
  const infList: Inf[] = [];
  for (const defKey in swg.definitions) {
    // console.debug(`formatDefs: ${defKey}`);
    const defItem = swg.definitions[defKey];
    const props: Prop[] = [];
    infList.push({ name: defKey, props });
    if (defItem.properties) {
      objectEntries(defItem.properties).forEach(([propName, prop]) => {
        props.push({ propName, propType: getType(prop) });
      });
    } else {
      LogDebug(`properties is null: ${defKey}`);
    }
  }
  let defs = "";
  infList.forEach((inf) => {
    const def =
      `interface ${inf.name} {\n` +
      inf.props.reduce((prev, curProp) => {
        return prev + `  ${curProp.propName}: ${curProp.propType}\n`;
      }, "") +
      "}\n";
    defs += def;
  });
  return defs;
}

function getParams(params: Swagger.Parameter[]) {
  return (params as Swagger.Query[])
    .filter((p) => p.in === "query")
    .map((p) => {
      return `${p.name}${p.required ? "" : "?"}:${getType(p)}`;
    });
}

function getBody(params: Swagger.Parameter[]) {
  return (params as Swagger.Body[])
    .filter((p) => p.in === "body")
    .map((p) => {
      return `${getType(p.schema)}`;
    });
}

function getType(prop?: Swagger.Query | Swagger.Schema | Swagger.Property) {
  if (prop) {
    if ("$ref" in prop) {
      return fromRefType(prop.$ref);
    } else if ("type" in prop) {
      if (prop.type === "array") {
        return fromArrayItems(prop.items);
      } else {
        return fromRawType(prop.type);
      }
    }
  }
  return "void";
}

// function isRefSchema(schema: Swagger.Schema): schema is Swagger.RefSchema {
//   return "$ref" in schema;
// }
// function isObjSchema(schema: Swagger.Schema): schema is Swagger.ObjSchema {
//   return "type" in schema && (schema as Swagger.ObjSchema).type === "object";
// }
// function isArraySchema(schema: Swagger.Schema): schema is Swagger.ArraySchema {
//   return "type" in schema && (schema as Swagger.ArraySchema).type === "array";
// }

function fromRefType(val: string) {
  const type = val.replace("#/definitions/", "");
  if (type === "Map«string,object»") return "{[k:string]:object}";
  return type;
}

function fromRawType(type: Swagger.RawType) {
  switch (type) {
    case "integer": {
      return "number";
    }
    default: {
      return type;
    }
  }
}

function fromArrayItems(items: Swagger.Items) {
  if ("$ref" in items) {
    return fromRefType(items.$ref) + "[]";
  } else {
    return fromRawType(items.type) + "[]";
  }
}
