// @generated by protobuf-ts 2.8.2
// @generated from protobuf file "trip.proto" (package "coolcar", syntax proto3)
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
 * @generated from protobuf message coolcar.Location
 */
export interface Location {
    /**
     * @generated from protobuf field: double latitude = 1;
     */
    latitude: number;
    /**
     * @generated from protobuf field: double longitude = 2;
     */
    longitude: number;
}
/**
 * @generated from protobuf message coolcar.Trip
 */
export interface Trip {
    /**
     * @generated from protobuf field: string start = 1;
     */
    start: string;
    /**
     * @generated from protobuf field: string end = 2;
     */
    end: string;
    /**
     * @generated from protobuf field: int32 duration_sec = 3;
     */
    durationSec: number;
    /**
     * @generated from protobuf field: int32 fee_cent = 4;
     */
    feeCent: number;
    /**
     * @generated from protobuf field: coolcar.Location start_pos = 5;
     */
    startPos?: Location;
    /**
     * @generated from protobuf field: coolcar.Location end_pos = 6;
     */
    endPos?: Location;
    /**
     * @generated from protobuf field: repeated coolcar.Location path_locations = 7;
     */
    pathLocations: Location[];
    /**
     * @generated from protobuf field: coolcar.TripStatus status = 8;
     */
    status: TripStatus;
}
/**
 * @generated from protobuf message coolcar.GetTripRequest
 */
export interface GetTripRequest {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
}
/**
 * @generated from protobuf message coolcar.GetTripResponse
 */
export interface GetTripResponse {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: coolcar.Trip trip = 2;
     */
    trip?: Trip;
}
/**
 * @generated from protobuf enum coolcar.TripStatus
 */
export enum TripStatus {
    /**
     * @generated from protobuf enum value: TS_NOT_SPECIFIED = 0;
     */
    TS_NOT_SPECIFIED = 0,
    /**
     * @generated from protobuf enum value: NOT_STSRTED = 1;
     */
    NOT_STSRTED = 1,
    /**
     * @generated from protobuf enum value: IN_PROGRESS = 2;
     */
    IN_PROGRESS = 2,
    /**
     * @generated from protobuf enum value: FINISHED = 3;
     */
    FINISHED = 3,
    /**
     * @generated from protobuf enum value: PAIED = 4;
     */
    PAIED = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class Location$Type extends MessageType<Location> {
    constructor() {
        super("coolcar.Location", [
            { no: 1, name: "latitude", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "longitude", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value?: PartialMessage<Location>): Location {
        const message = { latitude: 0, longitude: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Location>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Location): Location {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* double latitude */ 1:
                    message.latitude = reader.double();
                    break;
                case /* double longitude */ 2:
                    message.longitude = reader.double();
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
    internalBinaryWrite(message: Location, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* double latitude = 1; */
        if (message.latitude !== 0)
            writer.tag(1, WireType.Bit64).double(message.latitude);
        /* double longitude = 2; */
        if (message.longitude !== 0)
            writer.tag(2, WireType.Bit64).double(message.longitude);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.Location
 */
export const Location = new Location$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Trip$Type extends MessageType<Trip> {
    constructor() {
        super("coolcar.Trip", [
            { no: 1, name: "start", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "end", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "duration_sec", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "fee_cent", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "start_pos", kind: "message", T: () => Location },
            { no: 6, name: "end_pos", kind: "message", T: () => Location },
            { no: 7, name: "path_locations", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Location },
            { no: 8, name: "status", kind: "enum", T: () => ["coolcar.TripStatus", TripStatus] }
        ]);
    }
    create(value?: PartialMessage<Trip>): Trip {
        const message = { start: "", end: "", durationSec: 0, feeCent: 0, pathLocations: [], status: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Trip>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Trip): Trip {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string start */ 1:
                    message.start = reader.string();
                    break;
                case /* string end */ 2:
                    message.end = reader.string();
                    break;
                case /* int32 duration_sec */ 3:
                    message.durationSec = reader.int32();
                    break;
                case /* int32 fee_cent */ 4:
                    message.feeCent = reader.int32();
                    break;
                case /* coolcar.Location start_pos */ 5:
                    message.startPos = Location.internalBinaryRead(reader, reader.uint32(), options, message.startPos);
                    break;
                case /* coolcar.Location end_pos */ 6:
                    message.endPos = Location.internalBinaryRead(reader, reader.uint32(), options, message.endPos);
                    break;
                case /* repeated coolcar.Location path_locations */ 7:
                    message.pathLocations.push(Location.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* coolcar.TripStatus status */ 8:
                    message.status = reader.int32();
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
    internalBinaryWrite(message: Trip, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string start = 1; */
        if (message.start !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.start);
        /* string end = 2; */
        if (message.end !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.end);
        /* int32 duration_sec = 3; */
        if (message.durationSec !== 0)
            writer.tag(3, WireType.Varint).int32(message.durationSec);
        /* int32 fee_cent = 4; */
        if (message.feeCent !== 0)
            writer.tag(4, WireType.Varint).int32(message.feeCent);
        /* coolcar.Location start_pos = 5; */
        if (message.startPos)
            Location.internalBinaryWrite(message.startPos, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* coolcar.Location end_pos = 6; */
        if (message.endPos)
            Location.internalBinaryWrite(message.endPos, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* repeated coolcar.Location path_locations = 7; */
        for (let i = 0; i < message.pathLocations.length; i++)
            Location.internalBinaryWrite(message.pathLocations[i], writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* coolcar.TripStatus status = 8; */
        if (message.status !== 0)
            writer.tag(8, WireType.Varint).int32(message.status);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.Trip
 */
export const Trip = new Trip$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetTripRequest$Type extends MessageType<GetTripRequest> {
    constructor() {
        super("coolcar.GetTripRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<GetTripRequest>): GetTripRequest {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetTripRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetTripRequest): GetTripRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
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
    internalBinaryWrite(message: GetTripRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.GetTripRequest
 */
export const GetTripRequest = new GetTripRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetTripResponse$Type extends MessageType<GetTripResponse> {
    constructor() {
        super("coolcar.GetTripResponse", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "trip", kind: "message", T: () => Trip }
        ]);
    }
    create(value?: PartialMessage<GetTripResponse>): GetTripResponse {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetTripResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetTripResponse): GetTripResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* coolcar.Trip trip */ 2:
                    message.trip = Trip.internalBinaryRead(reader, reader.uint32(), options, message.trip);
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
    internalBinaryWrite(message: GetTripResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* coolcar.Trip trip = 2; */
        if (message.trip)
            Trip.internalBinaryWrite(message.trip, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.GetTripResponse
 */
export const GetTripResponse = new GetTripResponse$Type();
/**
 * @generated ServiceType for protobuf service coolcar.TripService
 */
export const TripService = new ServiceType("coolcar.TripService", [
    { name: "GetTrip", options: {}, I: GetTripRequest, O: GetTripResponse }
]);
