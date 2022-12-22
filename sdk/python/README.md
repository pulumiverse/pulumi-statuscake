# Statuscake Resource Provider

**This plugin requires pulumi cli v3.35.3 or later**

The Statuscake Resource Provider lets you manage [Statuscake](https://www.statuscake.com/) resources.

It is based on the statuscake terraform provider https://registry.terraform.io/providers/StatusCakeDev/statuscake/latest/docs

## Installing

If you get an error like this:

```
error: could not load plugin for statuscake provider 'urn:pulumi:my-stack::my-project::pulumi:providers:statuscake::my-provider': no resource plugin 'pulumi-resource-statuscake' found in the workspace at version vX.Y.Z or on your $PATH, install the plugin using `pulumi plugin install resource statuscake vX.Y.Z`
```

You can use this command to install the missing version:

```
pulumi plugin install resource statuscake vX.Y.Z --server github://api.github.com/pulumiverse
```

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/statuscake
```

or `yarn`:

```bash
yarn add @pulumiverse/statuscake
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse_statuscake
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-statuscake/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.Statuscake
```

## Configuration

The following configuration points are available for the `statuscake` provider:

- `statuscake:apiToken` (Required) This is the Statuscake API token. The value can be sourced from the `STATUSCAKE_API_TOKEN` environment variable.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/statuscake/api-docs/).
