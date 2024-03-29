// @generated by protobuf-ts 2.8.2
// @generated from protobuf file "auth.proto" (package "auth.v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message auth.v1.LoginRequest
 */
export interface LoginRequest {
    /**
     * @generated from protobuf field: string code = 1;
     */
    code: string;
}
/**
 * @generated from protobuf message auth.v1.LoginResponse
 */
export interface LoginResponse {
    /**
     * @generated from protobuf field: string access_token = 1;
     */
    accessToken: string;
    /**
     * @generated from protobuf field: int32 expires_in = 2;
     */
    expiresIn: number;
}
// @generated message type with reflection information, may provide speed optimized methods
class LoginRequest$Type extends MessageType<LoginRequest> {
    constructor() {
        super("auth.v1.LoginRequest", [
            { no: 1, name: "code", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<LoginRequest>): LoginRequest {
        const message = { code: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LoginRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LoginRequest): LoginRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string code */ 1:
                    message.code = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: LoginRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string code = 1; */
        if (message.code !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.code);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.v1.LoginRequest
 */
export const LoginRequest = new LoginRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class LoginResponse$Type extends MessageType<LoginResponse> {
    constructor() {
        super("auth.v1.LoginResponse", [
            { no: 1, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "expires_in", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<LoginResponse>): LoginResponse {
        const message = { accessToken: "", expiresIn: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LoginResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LoginResponse): LoginResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string access_token */ 1:
                    message.accessToken = reader.string();
                    break;
                case /* int32 expires_in */ 2:
                    message.expiresIn = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: LoginResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string access_token = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* int32 expires_in = 2; */
        if (message.expiresIn !== 0)
            writer.tag(2, WireType.Varint).int32(message.expiresIn);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.v1.LoginResponse
 */
export const LoginResponse = new LoginResponse$Type();
/**
 * @generated ServiceType for protobuf service auth.v1.AuthService
 */
export const AuthService = new ServiceType("auth.v1.AuthService", [
    { name: "Login", options: {}, I: LoginRequest, O: LoginResponse }
]);
