declare namespace Swagger {
  type RawType = "boolean" | "string" | /*  "array" | */ "integer" | "object";
  type Method = "get" | "delete" | "post";

  interface Root {
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
    additionalProperties: RefType | Type | TypeExt;
  }

  interface ArrayType {
    type: "array";
    items: Items;
  }

  type Items = Type | RefType | ArrayType;

  interface RefSchema extends RefType {}
  interface ArraySchema extends ArrayType {}
  interface ObjSchema extends ObjectType {}

  type Schema = ObjSchema | RefSchema | ArraySchema;

  interface TypeExt extends Type {
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

  interface ArrayQuery extends BaseQuery, ArrayType {
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

  type Definition = {
    type: string;
    properties: Properties;
  } | ObjectType

  interface Properties {
    [name: string]: Property;
  }

  type Property = Type | RefType | ArrayType | ObjectType;
}
