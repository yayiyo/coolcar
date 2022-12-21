"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TripService = exports.GetTripResponse = exports.GetTripRequest = exports.Trip = exports.Location = exports.TripStatus = void 0;
// @generated by protobuf-ts 2.8.2
// @generated from protobuf file "trip.proto" (package "coolcar", syntax proto3)
// tslint:disable
const runtime_rpc_1 = require("@protobuf-ts/runtime-rpc");
const runtime_1 = require("@protobuf-ts/runtime");
const runtime_2 = require("@protobuf-ts/runtime");
const runtime_3 = require("@protobuf-ts/runtime");
const runtime_4 = require("@protobuf-ts/runtime");
const runtime_5 = require("@protobuf-ts/runtime");
/**
 * @generated from protobuf enum coolcar.TripStatus
 */
var TripStatus;
(function (TripStatus) {
    /**
     * @generated from protobuf enum value: TS_NOT_SPECIFIED = 0;
     */
    TripStatus[TripStatus["TS_NOT_SPECIFIED"] = 0] = "TS_NOT_SPECIFIED";
    /**
     * @generated from protobuf enum value: NOT_STSRTED = 1;
     */
    TripStatus[TripStatus["NOT_STSRTED"] = 1] = "NOT_STSRTED";
    /**
     * @generated from protobuf enum value: IN_PROGRESS = 2;
     */
    TripStatus[TripStatus["IN_PROGRESS"] = 2] = "IN_PROGRESS";
    /**
     * @generated from protobuf enum value: FINISHED = 3;
     */
    TripStatus[TripStatus["FINISHED"] = 3] = "FINISHED";
    /**
     * @generated from protobuf enum value: PAIED = 4;
     */
    TripStatus[TripStatus["PAIED"] = 4] = "PAIED";
})(TripStatus = exports.TripStatus || (exports.TripStatus = {}));
// @generated message type with reflection information, may provide speed optimized methods
class Location$Type extends runtime_5.MessageType {
    constructor() {
        super("coolcar.Location", [
            { no: 1, name: "latitude", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "longitude", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { latitude: 0, longitude: 0 };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
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
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* double latitude = 1; */
        if (message.latitude !== 0)
            writer.tag(1, runtime_1.WireType.Bit64).double(message.latitude);
        /* double longitude = 2; */
        if (message.longitude !== 0)
            writer.tag(2, runtime_1.WireType.Bit64).double(message.longitude);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.Location
 */
exports.Location = new Location$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Trip$Type extends runtime_5.MessageType {
    constructor() {
        super("coolcar.Trip", [
            { no: 1, name: "start", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "end", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "duration_sec", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "fee_cent", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "start_pos", kind: "message", T: () => exports.Location },
            { no: 6, name: "end_pos", kind: "message", T: () => exports.Location },
            { no: 7, name: "path_locations", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => exports.Location },
            { no: 8, name: "status", kind: "enum", T: () => ["coolcar.TripStatus", TripStatus] }
        ]);
    }
    create(value) {
        const message = { start: "", end: "", durationSec: 0, feeCent: 0, pathLocations: [], status: 0 };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
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
                    message.startPos = exports.Location.internalBinaryRead(reader, reader.uint32(), options, message.startPos);
                    break;
                case /* coolcar.Location end_pos */ 6:
                    message.endPos = exports.Location.internalBinaryRead(reader, reader.uint32(), options, message.endPos);
                    break;
                case /* repeated coolcar.Location path_locations */ 7:
                    message.pathLocations.push(exports.Location.internalBinaryRead(reader, reader.uint32(), options));
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
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string start = 1; */
        if (message.start !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.start);
        /* string end = 2; */
        if (message.end !== "")
            writer.tag(2, runtime_1.WireType.LengthDelimited).string(message.end);
        /* int32 duration_sec = 3; */
        if (message.durationSec !== 0)
            writer.tag(3, runtime_1.WireType.Varint).int32(message.durationSec);
        /* int32 fee_cent = 4; */
        if (message.feeCent !== 0)
            writer.tag(4, runtime_1.WireType.Varint).int32(message.feeCent);
        /* coolcar.Location start_pos = 5; */
        if (message.startPos)
            exports.Location.internalBinaryWrite(message.startPos, writer.tag(5, runtime_1.WireType.LengthDelimited).fork(), options).join();
        /* coolcar.Location end_pos = 6; */
        if (message.endPos)
            exports.Location.internalBinaryWrite(message.endPos, writer.tag(6, runtime_1.WireType.LengthDelimited).fork(), options).join();
        /* repeated coolcar.Location path_locations = 7; */
        for (let i = 0; i < message.pathLocations.length; i++)
            exports.Location.internalBinaryWrite(message.pathLocations[i], writer.tag(7, runtime_1.WireType.LengthDelimited).fork(), options).join();
        /* coolcar.TripStatus status = 8; */
        if (message.status !== 0)
            writer.tag(8, runtime_1.WireType.Varint).int32(message.status);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.Trip
 */
exports.Trip = new Trip$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetTripRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("coolcar.GetTripRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
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
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.GetTripRequest
 */
exports.GetTripRequest = new GetTripRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetTripResponse$Type extends runtime_5.MessageType {
    constructor() {
        super("coolcar.GetTripResponse", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "trip", kind: "message", T: () => exports.Trip }
        ]);
    }
    create(value) {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* coolcar.Trip trip */ 2:
                    message.trip = exports.Trip.internalBinaryRead(reader, reader.uint32(), options, message.trip);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.id);
        /* coolcar.Trip trip = 2; */
        if (message.trip)
            exports.Trip.internalBinaryWrite(message.trip, writer.tag(2, runtime_1.WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message coolcar.GetTripResponse
 */
exports.GetTripResponse = new GetTripResponse$Type();
/**
 * @generated ServiceType for protobuf service coolcar.TripService
 */
exports.TripService = new runtime_rpc_1.ServiceType("coolcar.TripService", [
    { name: "GetTrip", options: {}, I: exports.GetTripRequest, O: exports.GetTripResponse }
]);
