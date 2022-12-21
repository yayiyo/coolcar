"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TripServiceClient = void 0;
const trip_1 = require("./trip");
const runtime_rpc_1 = require("@protobuf-ts/runtime-rpc");
/**
 * @generated from protobuf service coolcar.TripService
 */
class TripServiceClient {
    constructor(_transport) {
        this._transport = _transport;
        this.typeName = trip_1.TripService.typeName;
        this.methods = trip_1.TripService.methods;
        this.options = trip_1.TripService.options;
    }
    /**
     * @generated from protobuf rpc: GetTrip(coolcar.GetTripRequest) returns (coolcar.GetTripResponse);
     */
    getTrip(input, options) {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return (0, runtime_rpc_1.stackIntercept)("unary", this._transport, method, opt, input);
    }
}
exports.TripServiceClient = TripServiceClient;
