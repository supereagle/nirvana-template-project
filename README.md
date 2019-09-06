<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Nirvana Template Project](#nirvana-template-project)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Design](#design)
    - [Status](#status)
    - [See also](#see-also)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Nirvana Template Project

## About the project

The template is used to create golang project in [Nirvana](https://github.com/caicloud/nirvana). All nirvana projects must follow the conventions in the
template. Calling for exceptions must be brought up in the engineering team.

### API docs

The template doesn't have API docs. For web service, please include API docs here, whether it's
auto-generated or hand-written. For auto-generated API docs, you can also give instructions on the
build process.

### Design

The template follows project convention doc.

* [Repository Conventions](https://github.com/caicloud/engineering/blob/master/guidelines/repo_conventions.md)

### Status

The template project is in alpha status.

### See also

* [golang project template](https://github.com/caicloud/golang-template-project)
* [python project template](https://github.com/caicloud/python-template-project)
* [common project template](https://github.com/caicloud/common-template-project)

## Getting started

Below we describe the conventions or tools specific to golang project.

### Layout

```
.                                   #
├── Gopkg.toml                      #
├── Makefile                        #
├── OWNERS                          #
├── README.md                       #
├── apis                            # Store apidocs (swagger json)
├── bin                             # Store the compiled binary
├── build                           # Store Dockerfile
│   └── demo-admin                  #
│       └── Dockerfile              #
├── cmd                             # Store startup commands for project
│   └── demo-admin                  #
│       └── main.go                 #
├── docs                            # Store docs
│   └── README.md                   #
├── hack                            # Store scripts
│   ├── README.md                   #
│   ├── read_cpus_available.sh      # Script to read available cpus
│   └── script.sh                   #
├── nirvana.yaml                    # File to describes your project
├── pkg                             # Store structures and converters required by API, distinguish by version
│   ├── apis                        #
│   │   └── v1                      #
│   │       ├── converters          #
│   │       │   └── converters.go   #
│   │       └── types.go            #
│   ├── descriptors                 # Store API descriptions (routing and others), distinguish by version
│   │   └── v1                      #
│   │       ├── descriptors.go      #
│   │       └── message.go          # Store API definition of message
│   ├── filters                     # Store HTTP Request filter
│   │   └── filter.go               #
│   ├── handler                     # Store the logical processing required by APIs
│   │   └── message.go              #
│   ├── middlewares                 # Store middlewares
│   │   └── middlewares.go          #
│   ├── modifiers                   # Store definition modifiers
│   │   └── modifiers.go            #
│   └── version                     # Store version information of project
│       └── version.go              #
├── release                         #
│   └── template-admin.yaml         #
├── test                            # Store all tests (except unit tests), e.g. integration, e2e tests.
│   └── test_make.sh                #
└── vendor                          #
```

## Notes

* Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
* Every project **MUST** use `dep` for vendor management and **MUST** checkin `vendor` direcotry.
* `cmd` and `build` **MUST** have the same set of subdirectories for main targets
  * For example, `cmd/admin,cmd/controller` and `build/admin,build/controller`.
  * Dockerfile **MUST** be put under `build` directory even if you have only one Dockerfile.
