## Usage of private modules:

```export GOPRIVATE=github.wdf.sap.corp```

Note:
If we want to have more than one domain we can do the following:
```
export GOPRIVATE=github.wdf.sap.corp,some.git.domain.com
```

For fine-grained control over module download and validation, the GONOPROXY and GONOSUMDB environment variables accept the same kind of glob list and override GOPRIVATE for the specific decision of whether to use the proxy and checksum database, respectively.

NOTE:
Also it should work for docker files. IT IS NOT TESTED!

Useful links:
- https://golang.org/ref/mod#private-module-proxy-direct
- https://pkg.go.dev/cmd/go#hdr-Configuration_for_downloading_non_public_code
- https://stackoverflow.com/questions/65755940/go-get-fails-to-download-a-go-package-git-repository-hosted-on-a-github-entre
