// @generated by protobuf-ts 2.8.2
// @generated from protobuf file "rental.proto" (package "rental.v1", syntax proto3)
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
 * @generated from protobuf message rental.v1.Location
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
 * @generated from protobuf message rental.v1.LocationStatus
 */
export interface LocationStatus {
    /**
     * @generated from protobuf field: rental.v1.Location location = 1;
     */
    location?: Location;
    /**
     * @generated from protobuf field: int32 fee_cent = 2;
     */
    feeCent: number;
    /**
     * @generated from protobuf field: double km_driven = 3;
     */
    kmDriven: number;
    /**
     * @generated from protobuf field: string poi_name = 4;
     */
    poiName: string;
    /**
     * @generated from protobuf field: int64 timestamp_sec = 5;
     */
    timestampSec: bigint;
}
/**
 * @generated from protobuf message rental.v1.TripEntity
 */
export interface TripEntity {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: rental.v1.Trip trip = 2;
     */
    trip?: Trip;
}
/**
 * @generated from protobuf message rental.v1.Trip
 */
export interface Trip {
    /**
     * @generated from protobuf field: string account_id = 1;
     */
    accountId: string;
    /**
     * @generated from protobuf field: string car_id = 2;
     */
    carId: string;
    /**
     * @generated from protobuf field: rental.v1.LocationStatus start = 3;
     */
    start?: LocationStatus;
    /**
     * @generated from protobuf field: rental.v1.LocationStatus current = 4;
     */
    current?: LocationStatus;
    /**
     * @generated from protobuf field: rental.v1.LocationStatus end = 5;
     */
    end?: LocationStatus;
    /**
     * @generated from protobuf field: rental.v1.TripStatus status = 6;
     */
    status: TripStatus;
    /**
     * @generated from protobuf field: string identity_id = 7;
     */
    identityId: string;
}
/**
 * @generated from protobuf message rental.v1.CreateTripRequest
 */
export interface CreateTripRequest {
    /**
     * @generated from protobuf field: rental.v1.Location start = 1;
     */
    start?: Location;
    /**
     * @generated from protobuf field: string car_id = 2;
     */
    carId: string;
}
/**
 * @generated from protobuf message rental.v1.GetTripRequest
 */
export interface GetTripRequest {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
}
/**
 * @generated from protobuf message rental.v1.GetTripsRequest
 */
export interface GetTripsRequest {
    /**
     * @generated from protobuf field: rental.v1.TripStatus status = 1;
     */
    status: TripStatus;
}
/**
 * @generated from protobuf message rental.v1.GetTripsResponse
 */
export interface GetTripsResponse {
    /**
     * @generated from protobuf field: repeated rental.v1.TripEntity trips = 1;
     */
    trips: TripEntity[];
}
/**
 * @generated from protobuf message rental.v1.UpdateTripRequest
 */
export interface UpdateTripRequest {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: rental.v1.LocationStatus current = 2;
     */
    current?: LocationStatus;
    /**
     * @generated from protobuf field: bool end_trip = 3;
     */
    endTrip: boolean;
}
/**
 * @generated from protobuf enum rental.v1.TripStatus
 */
export enum TripStatus {
    /**
     * @generated from protobuf enum value: TS_NOT_SPECIFIED = 0;
     */
    TS_NOT_SPECIFIED = 0,
    /**
     * @generated from protobuf enum value: IN_PROGRESS = 1;
     */
    IN_PROGRESS = 1,
    /**
     * @generated from protobuf enum value: FINISHED = 2;
     */
    FINISHED = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class Location$Type extends MessageType<Location> {
    constructor() {
        super("rental.v1.Location", [
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
 * @generated MessageType for protobuf message rental.v1.Location
 */
export const Location = new Location$Type();
// @generated message type with reflection information, may provide speed optimized methods
class LocationStatus$Type extends MessageType<LocationStatus> {
    constructor() {
        super("rental.v1.LocationStatus", [
            { no: 1, name: "location", kind: "message", T: () => Location },
            { no: 2, name: "fee_cent", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "km_driven", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "poi_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "timestamp_sec", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<LocationStatus>): LocationStatus {
        const message = { feeCent: 0, kmDriven: 0, poiName: "", timestampSec: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LocationStatus>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LocationStatus): LocationStatus {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* rental.v1.Location location */ 1:
                    message.location = Location.internalBinaryRead(reader, reader.uint32(), options, message.location);
                    break;
                case /* int32 fee_cent */ 2:
                    message.feeCent = reader.int32();
                    break;
                case /* double km_driven */ 3:
                    message.kmDriven = reader.double();
                    break;
                case /* string poi_name */ 4:
                    message.poiName = reader.string();
                    break;
                case /* int64 timestamp_sec */ 5:
                    message.timestampSec = reader.int64().toBigInt();
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
    internalBinaryWrite(message: LocationStatus, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* rental.v1.Location location = 1; */
        if (message.location)
            Location.internalBinaryWrite(message.location, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* int32 fee_cent = 2; */
        if (message.feeCent !== 0)
            writer.tag(2, WireType.Varint).int32(message.feeCent);
        /* double km_driven = 3; */
        if (message.kmDriven !== 0)
            writer.tag(3, WireType.Bit64).double(message.kmDriven);
        /* string poi_name = 4; */
        if (message.poiName !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.poiName);
        /* int64 timestamp_sec = 5; */
        if (message.timestampSec !== 0n)
            writer.tag(5, WireType.Varint).int64(message.timestampSec);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message rental.v1.LocationStatus
 */
export const LocationStatus = new LocationStatus$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TripEntity$Type extends MessageType<TripEntity> {
    constructor() {
        super("rental.v1.TripEntity", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "trip", kind: "message", T: () => Trip }
        ]);
    }
    create(value?: PartialMessage<TripEntity>): TripEntity {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TripEntity>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TripEntity): TripEntity {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* rental.v1.Trip trip */ 2:
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
    internalBinaryWrite(message: TripEntity, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* rental.v1.Trip trip = 2; */
        if (message.trip)
            Trip.internalBinaryWrite(message.trip, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message rental.v1.TripEntity
 */
export const TripEntity = new TripEntity$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Trip$Type extends MessageType<Trip> {
    constructor() {
        super("rental.v1.Trip", [
            { no: 1, name: "account_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "car_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "start", kind: "message", T: () => LocationStatus },
            { no: 4, name: "current", kind: "message", T: () => LocationStatus },
            { no: 5, name: "end", kind: "message", T: () => LocationStatus },
            { no: 6, name: "status", kind: "enum", T: () => ["rental.v1.TripStatus", TripStatus] },
            { no: 7, name: "identity_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Trip>): Trip {
        const message = { accountId: "", carId: "", status: 0, identityId: "" };
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
                case /* string account_id */ 1:
                    message.accountId = reader.string();
                    break;
                case /* string car_id */ 2:
                    message.carId = reader.string();
                    break;
                case /* rental.v1.LocationStatus start */ 3:
                    message.start = LocationStatus.internalBinaryRead(reader, reader.uint32(), options, message.start);
                    break;
                case /* rental.v1.LocationStatus current */ 4:
                    message.current = LocationStatus.internalBinaryRead(reader, reader.uint32(), options, message.current);
                    break;
                case /* rental.v1.LocationStatus end */ 5:
                    message.end = LocationStatus.internalBinaryRead(reader, reader.uint32(), options, message.end);
                    break;
                case /* rental.v1.TripStatus status */ 6:
                    message.status = reader.int32();
                    break;
                case /* string identity_id */ 7:
                    message.identityId = reader.string();
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
        /* string account_id = 1; */
        if (message.accountId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accountId);
        /* string car_id = 2; */
        if (message.carId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.carId);
        /* rental.v1.LocationStatus start = 3; */
        if (message.start)
            LocationStatus.internalBinaryWrite(message.start, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* rental.v1.LocationStatus current = 4; */
        if (message.current)
            LocationStatus.internalBinaryWrite(message.current, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* rental.v1.LocationStatus end = 5; */
        if (message.end)
            LocationStatus.internalBinaryWrite(message.end, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* rental.v1.TripStatus status = 6; */
        if (message.status !== 0)
            writer.tag(6, WireType.Varint).int32(message.status);
        /* string identity_id = 7; */
        if (message.identityId !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.identityId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message rental.v1.Trip
 */
export const Trip = new Trip$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateTripRequest$Type extends MessageType<CreateTripRequest> {
    constructor() {
        super("rental.v1.CreateTripRequest", [
            { no: 1, name: "start", kind: "message", T: () => Location },
            { no: 2, name: "car_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<CreateTripRequest>): CreateTripRequest {
        const message = { carId: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<CreateTripRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateTripRequest): CreateTripRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* rental.v1.Location start */ 1:
                    message.start = Location.internalBinaryRead(reader, reader.uint32(), options, message.start);
                    break;
                case /* string car_id */ 2:
                    message.carId = reader.string();
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
    internalBinaryWrite(message: CreateTripRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* rental.v1.Location start = 1; */
        if (message.start)
            Location.internalBinaryWrite(message.start, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string car_id = 2; */
        if (message.carId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.carId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message rental.v1.CreateTripRequest
 */
export const CreateTripRequest = new CreateTripRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetTripRequest$Type extends MessageType<GetTripRequest> {
    constructor() {
        super("rental.v1.GetTripRequest", [
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
 * @generated MessageType for protobuf message rental.v1.GetTripRequest
 */
export const GetTripRequest = new GetTripRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetTripsRequest$Type extends MessageType<GetTripsRequest> {
    constructor() {
        super("rental.v1.GetTripsRequest", [
            { no: 1, name: "status", kind: "enum", T: () => ["rental.v1.TripStatus", TripStatus] }
        ]);
    }
    create(value?: PartialMessage<GetTripsRequest>): GetTripsRequest {
        const message = { status: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetTripsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetTripsRequest): GetTripsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* rental.v1.TripStatus status */ 1:
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
    internalBinaryWrite(message: GetTripsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* rental.v1.TripStatus status = 1; */
        if (message.status !== 0)
            writer.tag(1, WireType.Varint).int32(message.status);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message rental.v1.GetTripsRequest
 */
export const GetTripsRequest = new GetTripsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetTripsResponse$Type extends MessageType<GetTripsResponse> {
    constructor() {
        super("rental.v1.GetTripsResponse", [
            { no: 1, name: "trips", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => TripEntity }
        ]);
    }
    create(value?: PartialMessage<GetTripsResponse>): GetTripsResponse {
        const message = { trips: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetTripsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetTripsResponse): GetTripsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated rental.v1.TripEntity trips */ 1:
                    message.trips.push(TripEntity.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: GetTripsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated rental.v1.TripEntity trips = 1; */
        for (let i = 0; i < message.trips.length; i++)
            TripEntity.internalBinaryWrite(message.trips[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message rental.v1.GetTripsResponse
 */
export const GetTripsResponse = new GetTripsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateTripRequest$Type extends MessageType<UpdateTripRequest> {
    constructor() {
        super("rental.v1.UpdateTripRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "current", kind: "message", T: () => LocationStatus },
            { no: 3, name: "end_trip", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<UpdateTripRequest>): UpdateTripRequest {
        const message = { id: "", endTrip: false };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UpdateTripRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateTripRequest): UpdateTripRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* rental.v1.LocationStatus current */ 2:
                    message.current = LocationStatus.internalBinaryRead(reader, reader.uint32(), options, message.current);
                    break;
                case /* bool end_trip */ 3:
                    message.endTrip = reader.bool();
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
    internalBinaryWrite(message: UpdateTripRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* rental.v1.LocationStatus current = 2; */
        if (message.current)
            LocationStatus.internalBinaryWrite(message.current, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* bool end_trip = 3; */
        if (message.endTrip !== false)
            writer.tag(3, WireType.Varint).bool(message.endTrip);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message rental.v1.UpdateTripRequest
 */
export const UpdateTripRequest = new UpdateTripRequest$Type();
/**
 * @generated ServiceType for protobuf service rental.v1.TripService
 */
export const TripService = new ServiceType("rental.v1.TripService", [
    { name: "CreateTrip", options: {}, I: CreateTripRequest, O: TripEntity },
    { name: "GetTrip", options: {}, I: GetTripRequest, O: Trip },
    { name: "GetTrips", options: {}, I: GetTripsRequest, O: GetTripsResponse },
    { name: "UpdateTrip", options: {}, I: UpdateTripRequest, O: Trip }
]);
