// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("statuscake");

/**
 * The API token for operations. This can also be provided as an environment variable `STATUSCAKE_API_TOKEN`
 */
export declare const apiToken: string | undefined;
Object.defineProperty(exports, "apiToken", {
    get() {
        return __config.get("apiToken") ?? utilities.getEnv("STATUSCAKE_API_TOKEN");
    },
    enumerable: true,
});

/**
 * Maximum backoff period in seconds after failed API calls. This can also be provided as an environment variable
 * `STATUSCAKE_MAX_BACKOFF`
 */
export declare const maxBackoff: number | undefined;
Object.defineProperty(exports, "maxBackoff", {
    get() {
        return __config.getObject<number>("maxBackoff");
    },
    enumerable: true,
});

/**
 * Minimum backoff period in seconds after failed API calls. This can also be provided as an environment variable
 * `STATUSCAKE_MIN_BACKOFF`
 */
export declare const minBackoff: number | undefined;
Object.defineProperty(exports, "minBackoff", {
    get() {
        return __config.getObject<number>("minBackoff");
    },
    enumerable: true,
});

/**
 * Maximum number of retries to perform when an API request fails. This can also be provided as an environment variable
 * `STATUSCAKE_RETRIES`
 */
export declare const retries: number | undefined;
Object.defineProperty(exports, "retries", {
    get() {
        return __config.getObject<number>("retries");
    },
    enumerable: true,
});

/**
 * RPS limit to apply when making calls to the API. This can also be provided as an environment variable `STATUSCAKE_RPS`
 */
export declare const rps: number | undefined;
Object.defineProperty(exports, "rps", {
    get() {
        return __config.getObject<number>("rps");
    },
    enumerable: true,
});

/**
 * Custom endpoint to which request will be made. This can also be provided as an environment variable
 * `STATUCAKE_CUSTOM_ENDPOINT`
 */
export declare const statuscakeCustomEndpoint: string | undefined;
Object.defineProperty(exports, "statuscakeCustomEndpoint", {
    get() {
        return __config.get("statuscakeCustomEndpoint");
    },
    enumerable: true,
});
