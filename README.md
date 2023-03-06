# Library Golang VOIP Core

This is a basic VOIP core by [Golang](https://go.dev/)

### Documentation

- [Asterisk Manager Interface (AMI)](https://github.com/pnguyen215/gobase-voip-core/blob/master/docs/ami.md)

### Go Directories

<code>/cmd</code>
- This folder contains the main application entry point files for the project, with the directory name matching the name for the binary.
- Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!
- It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.

<code>/pkg</code>
- This folder contains code which is OK for other services to consume, this may include API clients, or utility functions which may be handy for other projects but don’t justify their own project.
- The pkg directory origins: The old Go source code used to use pkg for its packages and then various Go projects in the community started copying the pattern 

<code>/internal</code>
- This package holds the private library code used in your service, it is specific to the function of the service and not shared with other services.
- You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

<code>/vendor</code>
- Application dependencies (managed manually or by your favorite dependency management tool like the new built-in Go Modules feature). The go mod vendor command will create the /vendor directory for you. Note that you might need to add the -mod=vendor flag to your go build command if you are not using Go 1.14 where it's on by default.
- Don't commit your application dependencies if you are building a library.

<code>/api</code>
- OpenAPI/Swagger specs, JSON schema files, protocol definition files.

<code>/web</code>
- Web application specific components: static web assets, server side templates and SPAs.

<code>/configs</code>
- Configuration file templates or default configs. Put your `confd` or `consul-template` template files here.

<code>/init</code>
- System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

<code>/scripts</code>
- Scripts to perform various build, install, analysis, etc operations.
- These scripts keep the root level Makefile small and simple (e.g., https://github.com/hashicorp/terraform/blob/master/Makefile).
- example: https://github.com/golang-standards/project-layout/blob/master/scripts/README.md

<code>/build</code>
- Packaging and Continuous Integration. Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.
- Put your CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the `/build/ci` directory linking them to the location where the CI tools expect them (when possible).

<code>/deployments</code>
- IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh). Note that in some repos (especially apps deployed with kubernetes) this directory is called `/deploy`.

<code>/docs</code>
- Design and user documents (in addition to your godoc generated documentation).
- example: https://github.com/golang-standards/project-layout/blob/master/docs/README.md

<code>/tools</code>
- Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.
- example: https://github.com/golang-standards/project-layout/blob/master/tools/README.md

<code>/examples</code>
- Examples for your applications and/or public libraries.
- refs: https://github.com/golang-standards/project-layout/blob/master/examples/README.md

<code>/assets</code>
- Other assets to go along with your repository (images, logos, etc).

<code>go.mod</code>
- The `go.mod` file defines the module's module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.

<code>go.sum</code>
- The `go.sum` contains all the dependency check sums, and is managed by the go tools. The checksum present in go.sum file is used to validate the checksum of each of direct and indirect dependency to confirm that none of them has been modified.

### Naming Convention

- Package names should be lowercase. Don't use snake_case or camelCase.
- Avoid overly use terms like util, common, script etc
- Use singular form
```bash
fooApp/
  circle.yml
  Dockerfile
  cmd/
    foosrv/
      main.go
    foocli/
      main.go
  pkg/
    fs/
      fs.go
      fs_test.go
      mock.go
      mock_test.go
    merge/
      merge.go
      merge_test.go
    api/
      api.go
      api_test.go
```

> As you noticed, there is func_test.go file in the same directory. In Go, you save unit tests inside separate files with a filename ending with _test.go. Go provides go test command out of the box which executes these files and runs tests.


### Example

#### Go Directories

```
├── LICENSE
├── README.md
├── config.go
├── go.mod
├── go.sum
├── client-lib
│   ├── lib.go
│   └── lib_test.go
├── cmd
│   ├── mod-lib-client
│   │   └── main.go
│   └── mod-lib-server
│       └── main.go
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
└── server-lib
    └── lib.go
```

```
$ tree exitus/
 exitus/
├── cmd
│   ├── authtest
│   │   └── main.go
│   ├── backend
│   │   └── main.go
│   └── client
│       └── main.go
├── dev
│   ├── add_migration.sh
│   └── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
│   ├── 20190721131113_extensions.down.sql
│   ├── 20190721131113_extensions.up.sql
│   ├── 20190723044115_customer_projects.down.sql
│   ├── 20190723044115_customer_projects.up.sql
│   ├── 20190726175158_issues.down.sql
│   ├── 20190726175158_issues.up.sql
│   ├── 20190726201649_comments.down.sql
│   ├── 20190726201649_comments.up.sql
│   ├── bindata.go
│   ├── gen.go
│   ├── migrations_test.go
│   └── README.md
├── pkg
│   ├── api
│   │   ├── exitus.gen.go
│   │   ├── exitus.yml
│   │   └── gen.go
│   ├── auth
│   │   ├── scopes.go
│   │   └── user.go
│   ├── conf
│   │   ├── conf.go
│   │   └── conf_test.go
│   ├── db
│   │   ├── db.go
│   │   ├── dbtesting.go
│   │   ├── migrate.go
│   │   ├── sqlhooks.go
│   │   └── transactions.go
│   ├── env
│   │   └── env.go
│   ├── healthz
│   │   ├── healthz.go
│   │   └── healthz_test.go
│   ├── jwt
│   │   └── jwt.go
│   ├── metrics
│   │   └── metrics.go
│   ├── middleware
│   │   ├── jwt.go
│   │   └── middleware.go
│   ├── oidc
│   │   └── client.go
│   ├── server
│   │   ├── reflect.go
│   │   └── server.go
│   └── store
│       ├── comments.go
│       ├── comments_test.go
│       ├── customers.go
│       ├── customers_test.go
│       ├── issues.go
│       ├── issues_test.go
│       ├── migrate_test.go
│       ├── projects.go
│       ├── projects_test.go
│       └── store.go
└── README.md
```

#### Reference
* https://github.com/golang-standards/project-layout
* https://dev.to/jinxankit/go-project-structure-and-guidelines-4ccm
* https://www.developer.com/languages/go-project-layout/
* https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/
* https://github.com/onmyway133/awesome-voip/blob/master/README.md