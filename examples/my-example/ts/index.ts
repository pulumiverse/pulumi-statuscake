import * as statuscake from "@pulumiverse/statuscake";

export const uptimeCheck = new statuscake.UptimeCheck("uptime", {
  checkInterval: 60,
  monitoredResource: { address: "https://www.pulumi.com" },
  httpCheck: { followRedirects: true, statusCodes: ["200"], validateSsl: true },
});
