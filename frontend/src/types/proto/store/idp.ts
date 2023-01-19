/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "bytebase.store";

export enum IdentityProviderType {
  IDENTITY_PROVIDER_TYPE_UNSPECIFIED = 0,
  OAUTH2 = 1,
  OIDC = 2,
  UNRECOGNIZED = -1,
}

export function identityProviderTypeFromJSON(object: any): IdentityProviderType {
  switch (object) {
    case 0:
    case "IDENTITY_PROVIDER_TYPE_UNSPECIFIED":
      return IdentityProviderType.IDENTITY_PROVIDER_TYPE_UNSPECIFIED;
    case 1:
    case "OAUTH2":
      return IdentityProviderType.OAUTH2;
    case 2:
    case "OIDC":
      return IdentityProviderType.OIDC;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IdentityProviderType.UNRECOGNIZED;
  }
}

export function identityProviderTypeToJSON(object: IdentityProviderType): string {
  switch (object) {
    case IdentityProviderType.IDENTITY_PROVIDER_TYPE_UNSPECIFIED:
      return "IDENTITY_PROVIDER_TYPE_UNSPECIFIED";
    case IdentityProviderType.OAUTH2:
      return "OAUTH2";
    case IdentityProviderType.OIDC:
      return "OIDC";
    case IdentityProviderType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface IdentityProviderConfig {
  oauth2Config?: OAuth2IdentityProviderConfig | undefined;
  oidcConfig?: OIDCIdentityProviderConfig | undefined;
}

/** OAuth2IdentityProviderConfig is the structure for OAuth2 identity provider config. */
export interface OAuth2IdentityProviderConfig {
  authUrl: string;
  tokenUrl: string;
  userInfoUrl: string;
  clientId: string;
  clientSecret: string;
  scopes: string[];
  fieldMapping?: FieldMapping;
}

/** OIDCIdentityProviderConfig is the structure for OIDC identity provider config. */
export interface OIDCIdentityProviderConfig {
  issuer: string;
  clientId: string;
  clientSecret: string;
  fieldMapping?: FieldMapping;
}

/** FieldMapping uses for mapping user info from identity provider to Bytebase. */
export interface FieldMapping {
  /** Identifier is the field name of the unique identifier in 3rd-party idp user info. Required. */
  identifier: string;
  /** Username is the field name of username in 3rd-party idp user info. Required. */
  username: string;
  /** Email is the field name of primary email in 3rd-party idp user info. Required. */
  email: string;
}

function createBaseIdentityProviderConfig(): IdentityProviderConfig {
  return { oauth2Config: undefined, oidcConfig: undefined };
}

export const IdentityProviderConfig = {
  encode(message: IdentityProviderConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.oauth2Config !== undefined) {
      OAuth2IdentityProviderConfig.encode(message.oauth2Config, writer.uint32(10).fork()).ldelim();
    }
    if (message.oidcConfig !== undefined) {
      OIDCIdentityProviderConfig.encode(message.oidcConfig, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IdentityProviderConfig {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.oauth2Config = OAuth2IdentityProviderConfig.decode(reader, reader.uint32());
          break;
        case 2:
          message.oidcConfig = OIDCIdentityProviderConfig.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IdentityProviderConfig {
    return {
      oauth2Config: isSet(object.oauth2Config) ? OAuth2IdentityProviderConfig.fromJSON(object.oauth2Config) : undefined,
      oidcConfig: isSet(object.oidcConfig) ? OIDCIdentityProviderConfig.fromJSON(object.oidcConfig) : undefined,
    };
  },

  toJSON(message: IdentityProviderConfig): unknown {
    const obj: any = {};
    message.oauth2Config !== undefined &&
      (obj.oauth2Config = message.oauth2Config ? OAuth2IdentityProviderConfig.toJSON(message.oauth2Config) : undefined);
    message.oidcConfig !== undefined &&
      (obj.oidcConfig = message.oidcConfig ? OIDCIdentityProviderConfig.toJSON(message.oidcConfig) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<IdentityProviderConfig>): IdentityProviderConfig {
    const message = createBaseIdentityProviderConfig();
    message.oauth2Config = (object.oauth2Config !== undefined && object.oauth2Config !== null)
      ? OAuth2IdentityProviderConfig.fromPartial(object.oauth2Config)
      : undefined;
    message.oidcConfig = (object.oidcConfig !== undefined && object.oidcConfig !== null)
      ? OIDCIdentityProviderConfig.fromPartial(object.oidcConfig)
      : undefined;
    return message;
  },
};

function createBaseOAuth2IdentityProviderConfig(): OAuth2IdentityProviderConfig {
  return {
    authUrl: "",
    tokenUrl: "",
    userInfoUrl: "",
    clientId: "",
    clientSecret: "",
    scopes: [],
    fieldMapping: undefined,
  };
}

export const OAuth2IdentityProviderConfig = {
  encode(message: OAuth2IdentityProviderConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authUrl !== "") {
      writer.uint32(10).string(message.authUrl);
    }
    if (message.tokenUrl !== "") {
      writer.uint32(18).string(message.tokenUrl);
    }
    if (message.userInfoUrl !== "") {
      writer.uint32(26).string(message.userInfoUrl);
    }
    if (message.clientId !== "") {
      writer.uint32(34).string(message.clientId);
    }
    if (message.clientSecret !== "") {
      writer.uint32(42).string(message.clientSecret);
    }
    for (const v of message.scopes) {
      writer.uint32(50).string(v!);
    }
    if (message.fieldMapping !== undefined) {
      FieldMapping.encode(message.fieldMapping, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OAuth2IdentityProviderConfig {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOAuth2IdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authUrl = reader.string();
          break;
        case 2:
          message.tokenUrl = reader.string();
          break;
        case 3:
          message.userInfoUrl = reader.string();
          break;
        case 4:
          message.clientId = reader.string();
          break;
        case 5:
          message.clientSecret = reader.string();
          break;
        case 6:
          message.scopes.push(reader.string());
          break;
        case 7:
          message.fieldMapping = FieldMapping.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OAuth2IdentityProviderConfig {
    return {
      authUrl: isSet(object.authUrl) ? String(object.authUrl) : "",
      tokenUrl: isSet(object.tokenUrl) ? String(object.tokenUrl) : "",
      userInfoUrl: isSet(object.userInfoUrl) ? String(object.userInfoUrl) : "",
      clientId: isSet(object.clientId) ? String(object.clientId) : "",
      clientSecret: isSet(object.clientSecret) ? String(object.clientSecret) : "",
      scopes: Array.isArray(object?.scopes) ? object.scopes.map((e: any) => String(e)) : [],
      fieldMapping: isSet(object.fieldMapping) ? FieldMapping.fromJSON(object.fieldMapping) : undefined,
    };
  },

  toJSON(message: OAuth2IdentityProviderConfig): unknown {
    const obj: any = {};
    message.authUrl !== undefined && (obj.authUrl = message.authUrl);
    message.tokenUrl !== undefined && (obj.tokenUrl = message.tokenUrl);
    message.userInfoUrl !== undefined && (obj.userInfoUrl = message.userInfoUrl);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.clientSecret !== undefined && (obj.clientSecret = message.clientSecret);
    if (message.scopes) {
      obj.scopes = message.scopes.map((e) => e);
    } else {
      obj.scopes = [];
    }
    message.fieldMapping !== undefined &&
      (obj.fieldMapping = message.fieldMapping ? FieldMapping.toJSON(message.fieldMapping) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<OAuth2IdentityProviderConfig>): OAuth2IdentityProviderConfig {
    const message = createBaseOAuth2IdentityProviderConfig();
    message.authUrl = object.authUrl ?? "";
    message.tokenUrl = object.tokenUrl ?? "";
    message.userInfoUrl = object.userInfoUrl ?? "";
    message.clientId = object.clientId ?? "";
    message.clientSecret = object.clientSecret ?? "";
    message.scopes = object.scopes?.map((e) => e) || [];
    message.fieldMapping = (object.fieldMapping !== undefined && object.fieldMapping !== null)
      ? FieldMapping.fromPartial(object.fieldMapping)
      : undefined;
    return message;
  },
};

function createBaseOIDCIdentityProviderConfig(): OIDCIdentityProviderConfig {
  return { issuer: "", clientId: "", clientSecret: "", fieldMapping: undefined };
}

export const OIDCIdentityProviderConfig = {
  encode(message: OIDCIdentityProviderConfig, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    if (message.clientId !== "") {
      writer.uint32(18).string(message.clientId);
    }
    if (message.clientSecret !== "") {
      writer.uint32(26).string(message.clientSecret);
    }
    if (message.fieldMapping !== undefined) {
      FieldMapping.encode(message.fieldMapping, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OIDCIdentityProviderConfig {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOIDCIdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string();
          break;
        case 2:
          message.clientId = reader.string();
          break;
        case 3:
          message.clientSecret = reader.string();
          break;
        case 4:
          message.fieldMapping = FieldMapping.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OIDCIdentityProviderConfig {
    return {
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      clientId: isSet(object.clientId) ? String(object.clientId) : "",
      clientSecret: isSet(object.clientSecret) ? String(object.clientSecret) : "",
      fieldMapping: isSet(object.fieldMapping) ? FieldMapping.fromJSON(object.fieldMapping) : undefined,
    };
  },

  toJSON(message: OIDCIdentityProviderConfig): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.clientSecret !== undefined && (obj.clientSecret = message.clientSecret);
    message.fieldMapping !== undefined &&
      (obj.fieldMapping = message.fieldMapping ? FieldMapping.toJSON(message.fieldMapping) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<OIDCIdentityProviderConfig>): OIDCIdentityProviderConfig {
    const message = createBaseOIDCIdentityProviderConfig();
    message.issuer = object.issuer ?? "";
    message.clientId = object.clientId ?? "";
    message.clientSecret = object.clientSecret ?? "";
    message.fieldMapping = (object.fieldMapping !== undefined && object.fieldMapping !== null)
      ? FieldMapping.fromPartial(object.fieldMapping)
      : undefined;
    return message;
  },
};

function createBaseFieldMapping(): FieldMapping {
  return { identifier: "", username: "", email: "" };
}

export const FieldMapping = {
  encode(message: FieldMapping, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.identifier !== "") {
      writer.uint32(10).string(message.identifier);
    }
    if (message.username !== "") {
      writer.uint32(18).string(message.username);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FieldMapping {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFieldMapping();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.identifier = reader.string();
          break;
        case 2:
          message.username = reader.string();
          break;
        case 3:
          message.email = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FieldMapping {
    return {
      identifier: isSet(object.identifier) ? String(object.identifier) : "",
      username: isSet(object.username) ? String(object.username) : "",
      email: isSet(object.email) ? String(object.email) : "",
    };
  },

  toJSON(message: FieldMapping): unknown {
    const obj: any = {};
    message.identifier !== undefined && (obj.identifier = message.identifier);
    message.username !== undefined && (obj.username = message.username);
    message.email !== undefined && (obj.email = message.email);
    return obj;
  },

  fromPartial(object: DeepPartial<FieldMapping>): FieldMapping {
    const message = createBaseFieldMapping();
    message.identifier = object.identifier ?? "";
    message.username = object.username ?? "";
    message.email = object.email ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
