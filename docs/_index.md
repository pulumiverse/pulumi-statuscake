---
title: Statuscake
meta_desc: Provides an overview of the Statuscake Provider for Pulumi.
layout: overview
---

The Statuscake provider for Pulumi can be used to provision resources in [Statuscake](https://www.statuscake.com).

The Statuscake provider must be configured with credentials to create and update resources in Statuscake.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as statuscake from "@pulumiverse/statuscake";

export const uptimeCheck = new statuscake.StatuscakeUptimeCheck("example", {
  checkInterval: 60,
  monitoredResource: { address: "https://www.pulumi.com" },
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_statuscake as statuscake

uptimeCheck = statuscake.StatuscakeUptimeCheck("example",
    checkInterval=60
    monitoredResource={
        address="https://www.pulumi.com"
    }
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	statuscake "github.com/pulumiverse/pulumi-statuscake/sdk/go/statuscake"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := statuscake.StatuscakeUptimeCheck(ctx, "example", &statuscake.StatuscakeUptimeCheckArgs{
            CheckInterval: 60,
            MonitoredResource: statuscake.StatuscakeUptimeCheckMonitoredResource{
                Address: "https://www.pulumi.com",
            },
		})

		if err != nil {
			return fmt.Errorf("error creating statuscake uptime check: %v", err)
		}

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.statuscake;

class statusCakeUptimeCheck : Stack
{
    public uptimeCheck()
    {
        var project = new StatuscakeUptimeCheck("example", new StatuscakeUptimeCheckArgs{
            CheckInterval: 60,
            MonitoredResource: {
                Address: "https://www.pulumi.com",
            }
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
