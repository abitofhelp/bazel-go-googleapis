#bazel-go-googleapis.git

This repository contains a client/server solution implementing the classic Greet service using the following technologies:
* Bazel
* Gazelle
* Go
* GoogleApis (Bazel Central Registry)
* Grpc
* Grpc Reflection
* Protocol Buffers

My goal was to explore how to use the googleapis proto package from BCR rather than having to 
configure the following in the MODULE.bazel file.

```
# See:
# https://bazel.build/external/migration#fetch-deps-module-extensions
# https://github.com/bazelbuild/rules_go/issues/3685

GOOGLE_APIS_VERSION = "64926d52febbf298cb82a8f472ade4a3969ba922"

bazel_dep(
    name = "com_google_googleapis",
    version = GOOGLE_APIS_VERSION,
    repo_name = "googleapis",
)
archive_override(
    module_name = "com_google_googleapis",
    integrity = "sha256-nRqTDnZ8k8glOYuPhpLso/41O5qq3t+88fyiKCyF34g=",
    patch_strip = 1,
    patches = [
        "scripts/bazel/add_bzlmod_support.patch",  # https://github.com/googleapis/googleapis/pull/855
    ],
    strip_prefix = "googleapis-" + GOOGLE_APIS_VERSION,
    urls = [
        "https://github.com/googleapis/googleapis/archive/%s.zip" % GOOGLE_APIS_VERSION,
    ],
)
```