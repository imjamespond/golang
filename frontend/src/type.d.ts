declare namespace Swagger {
  type RawType = "boolean" | "string" | /*  "array" | */ "integer" | "object";
  type Method = "get" | "delete" | "post";

  interface RootObject {
    swagger: string;
    info: Info;
    host: string;
    basePath: string;
    tags: Tag[];
    paths: Paths;
    definitions: Definitions;
  }

  interface Definitions {
    [name: string]: Definition;
  }

  type Paths = Record<string, Path>;

  type Path = Record<Method, PathItem>;

  interface PathItem {
    tags: string[];
    summary: string;
    description: string;
    operationId: string;
    consumes: string[];
    produces: string[];
    parameters?: Parameter[];
    responses: Response.Responses;
  }

  interface RefType {
    $ref: string;
  }

  interface Type {
    type: RawType;
  }

  interface ObjectType {
    type: "object";
    additionalProperties: RefType | Type | AdditionalProperties;
  }

  interface ArrayType {
    type: "array";
    items: Items;
  }

  type Items = Type | RefType;

  interface RefSchema extends RefType {}
  interface ArraySchema extends ArrayType {}
  interface ObjSchema extends ObjectType {}

  type Schema = ObjSchema | RefSchema | ArraySchema;

  interface AdditionalProperties extends Type {
    format: string;
  }

  interface BaseQuery {
    name: string;
    in: "query";
    description: string;
    required: boolean;
  }

  interface SimpleQuery extends BaseQuery {
    type: RawType;
  }

  interface ArrayQuery extends BaseQuery {
    type: "array";
    items: Items;
    collectionFormat: string;
  }

  type Query = SimpleQuery | ArrayQuery;

  interface Body {
    in: "body";
    name: string;
    description: string;
    required: boolean;
    schema: Schema;
  }

  type Parameter = Query | Body;

  declare namespace Response {
    interface Responses {
      "200": Response;
      "204": Response;
      "401": Response;
      "403": Response;
    }

    interface Response {
      description: string;
      schema?: Schema;
    }
  }

  interface Definition {
    type: string;
    properties: Properties;
  }

  interface Properties {
    [name: string]: Property;
  }

  type Property = Type | RefType | ArrayType | ObjectType;
}
