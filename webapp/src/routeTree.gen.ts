/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as AboutImport } from './routes/about'
import { Route as IndexImport } from './routes/index'
import { Route as ServerGroupsOldindexImport } from './routes/server-groups/oldindex'
import { Route as TestsIndexImport } from './routes/tests/index'
import { Route as ServerPortIndexImport } from './routes/server-port/index'
import { Route as ServerGroupsIndexImport } from './routes/server-groups/index'
import { Route as ServerPortStatusImport } from './routes/server-port/status'
import { Route as ServerPortFormImport } from './routes/server-port/form'
import { Route as ServerGroupsIncidentsImport } from './routes/server-groups/incidents'
import { Route as ServerGroupsAddServerImport } from './routes/server-groups/add-server'
import { Route as ServerGroupsAddImport } from './routes/server-groups/add'
import { Route as ConfigNatsIndexImport } from './routes/config/nats/index'
import { Route as ServerGroupsManageServersGroupIdImport } from './routes/server-groups/manage-servers.$groupId'

// Create/Update Routes

const AboutRoute = AboutImport.update({
  id: '/about',
  path: '/about',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const ServerGroupsOldindexRoute = ServerGroupsOldindexImport.update({
  id: '/server-groups/oldindex',
  path: '/server-groups/oldindex',
  getParentRoute: () => rootRoute,
} as any)

const TestsIndexRoute = TestsIndexImport.update({
  id: '/tests/',
  path: '/tests/',
  getParentRoute: () => rootRoute,
} as any)

const ServerPortIndexRoute = ServerPortIndexImport.update({
  id: '/server-port/',
  path: '/server-port/',
  getParentRoute: () => rootRoute,
} as any)

const ServerGroupsIndexRoute = ServerGroupsIndexImport.update({
  id: '/server-groups/',
  path: '/server-groups/',
  getParentRoute: () => rootRoute,
} as any)

const ServerPortStatusRoute = ServerPortStatusImport.update({
  id: '/server-port/status',
  path: '/server-port/status',
  getParentRoute: () => rootRoute,
} as any)

const ServerPortFormRoute = ServerPortFormImport.update({
  id: '/server-port/form',
  path: '/server-port/form',
  getParentRoute: () => rootRoute,
} as any)

const ServerGroupsIncidentsRoute = ServerGroupsIncidentsImport.update({
  id: '/server-groups/incidents',
  path: '/server-groups/incidents',
  getParentRoute: () => rootRoute,
} as any)

const ServerGroupsAddServerRoute = ServerGroupsAddServerImport.update({
  id: '/server-groups/add-server',
  path: '/server-groups/add-server',
  getParentRoute: () => rootRoute,
} as any)

const ServerGroupsAddRoute = ServerGroupsAddImport.update({
  id: '/server-groups/add',
  path: '/server-groups/add',
  getParentRoute: () => rootRoute,
} as any)

const ConfigNatsIndexRoute = ConfigNatsIndexImport.update({
  id: '/config/nats/',
  path: '/config/nats/',
  getParentRoute: () => rootRoute,
} as any)

const ServerGroupsManageServersGroupIdRoute =
  ServerGroupsManageServersGroupIdImport.update({
    id: '/server-groups/manage-servers/$groupId',
    path: '/server-groups/manage-servers/$groupId',
    getParentRoute: () => rootRoute,
  } as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutImport
      parentRoute: typeof rootRoute
    }
    '/server-groups/add': {
      id: '/server-groups/add'
      path: '/server-groups/add'
      fullPath: '/server-groups/add'
      preLoaderRoute: typeof ServerGroupsAddImport
      parentRoute: typeof rootRoute
    }
    '/server-groups/add-server': {
      id: '/server-groups/add-server'
      path: '/server-groups/add-server'
      fullPath: '/server-groups/add-server'
      preLoaderRoute: typeof ServerGroupsAddServerImport
      parentRoute: typeof rootRoute
    }
    '/server-groups/incidents': {
      id: '/server-groups/incidents'
      path: '/server-groups/incidents'
      fullPath: '/server-groups/incidents'
      preLoaderRoute: typeof ServerGroupsIncidentsImport
      parentRoute: typeof rootRoute
    }
    '/server-groups/oldindex': {
      id: '/server-groups/oldindex'
      path: '/server-groups/oldindex'
      fullPath: '/server-groups/oldindex'
      preLoaderRoute: typeof ServerGroupsOldindexImport
      parentRoute: typeof rootRoute
    }
    '/server-port/form': {
      id: '/server-port/form'
      path: '/server-port/form'
      fullPath: '/server-port/form'
      preLoaderRoute: typeof ServerPortFormImport
      parentRoute: typeof rootRoute
    }
    '/server-port/status': {
      id: '/server-port/status'
      path: '/server-port/status'
      fullPath: '/server-port/status'
      preLoaderRoute: typeof ServerPortStatusImport
      parentRoute: typeof rootRoute
    }
    '/server-groups/': {
      id: '/server-groups/'
      path: '/server-groups'
      fullPath: '/server-groups'
      preLoaderRoute: typeof ServerGroupsIndexImport
      parentRoute: typeof rootRoute
    }
    '/server-port/': {
      id: '/server-port/'
      path: '/server-port'
      fullPath: '/server-port'
      preLoaderRoute: typeof ServerPortIndexImport
      parentRoute: typeof rootRoute
    }
    '/tests/': {
      id: '/tests/'
      path: '/tests'
      fullPath: '/tests'
      preLoaderRoute: typeof TestsIndexImport
      parentRoute: typeof rootRoute
    }
    '/server-groups/manage-servers/$groupId': {
      id: '/server-groups/manage-servers/$groupId'
      path: '/server-groups/manage-servers/$groupId'
      fullPath: '/server-groups/manage-servers/$groupId'
      preLoaderRoute: typeof ServerGroupsManageServersGroupIdImport
      parentRoute: typeof rootRoute
    }
    '/config/nats/': {
      id: '/config/nats/'
      path: '/config/nats'
      fullPath: '/config/nats'
      preLoaderRoute: typeof ConfigNatsIndexImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexRoute
  '/about': typeof AboutRoute
  '/server-groups/add': typeof ServerGroupsAddRoute
  '/server-groups/add-server': typeof ServerGroupsAddServerRoute
  '/server-groups/incidents': typeof ServerGroupsIncidentsRoute
  '/server-groups/oldindex': typeof ServerGroupsOldindexRoute
  '/server-port/form': typeof ServerPortFormRoute
  '/server-port/status': typeof ServerPortStatusRoute
  '/server-groups': typeof ServerGroupsIndexRoute
  '/server-port': typeof ServerPortIndexRoute
  '/tests': typeof TestsIndexRoute
  '/server-groups/manage-servers/$groupId': typeof ServerGroupsManageServersGroupIdRoute
  '/config/nats': typeof ConfigNatsIndexRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexRoute
  '/about': typeof AboutRoute
  '/server-groups/add': typeof ServerGroupsAddRoute
  '/server-groups/add-server': typeof ServerGroupsAddServerRoute
  '/server-groups/incidents': typeof ServerGroupsIncidentsRoute
  '/server-groups/oldindex': typeof ServerGroupsOldindexRoute
  '/server-port/form': typeof ServerPortFormRoute
  '/server-port/status': typeof ServerPortStatusRoute
  '/server-groups': typeof ServerGroupsIndexRoute
  '/server-port': typeof ServerPortIndexRoute
  '/tests': typeof TestsIndexRoute
  '/server-groups/manage-servers/$groupId': typeof ServerGroupsManageServersGroupIdRoute
  '/config/nats': typeof ConfigNatsIndexRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexRoute
  '/about': typeof AboutRoute
  '/server-groups/add': typeof ServerGroupsAddRoute
  '/server-groups/add-server': typeof ServerGroupsAddServerRoute
  '/server-groups/incidents': typeof ServerGroupsIncidentsRoute
  '/server-groups/oldindex': typeof ServerGroupsOldindexRoute
  '/server-port/form': typeof ServerPortFormRoute
  '/server-port/status': typeof ServerPortStatusRoute
  '/server-groups/': typeof ServerGroupsIndexRoute
  '/server-port/': typeof ServerPortIndexRoute
  '/tests/': typeof TestsIndexRoute
  '/server-groups/manage-servers/$groupId': typeof ServerGroupsManageServersGroupIdRoute
  '/config/nats/': typeof ConfigNatsIndexRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths:
    | '/'
    | '/about'
    | '/server-groups/add'
    | '/server-groups/add-server'
    | '/server-groups/incidents'
    | '/server-groups/oldindex'
    | '/server-port/form'
    | '/server-port/status'
    | '/server-groups'
    | '/server-port'
    | '/tests'
    | '/server-groups/manage-servers/$groupId'
    | '/config/nats'
  fileRoutesByTo: FileRoutesByTo
  to:
    | '/'
    | '/about'
    | '/server-groups/add'
    | '/server-groups/add-server'
    | '/server-groups/incidents'
    | '/server-groups/oldindex'
    | '/server-port/form'
    | '/server-port/status'
    | '/server-groups'
    | '/server-port'
    | '/tests'
    | '/server-groups/manage-servers/$groupId'
    | '/config/nats'
  id:
    | '__root__'
    | '/'
    | '/about'
    | '/server-groups/add'
    | '/server-groups/add-server'
    | '/server-groups/incidents'
    | '/server-groups/oldindex'
    | '/server-port/form'
    | '/server-port/status'
    | '/server-groups/'
    | '/server-port/'
    | '/tests/'
    | '/server-groups/manage-servers/$groupId'
    | '/config/nats/'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexRoute: typeof IndexRoute
  AboutRoute: typeof AboutRoute
  ServerGroupsAddRoute: typeof ServerGroupsAddRoute
  ServerGroupsAddServerRoute: typeof ServerGroupsAddServerRoute
  ServerGroupsIncidentsRoute: typeof ServerGroupsIncidentsRoute
  ServerGroupsOldindexRoute: typeof ServerGroupsOldindexRoute
  ServerPortFormRoute: typeof ServerPortFormRoute
  ServerPortStatusRoute: typeof ServerPortStatusRoute
  ServerGroupsIndexRoute: typeof ServerGroupsIndexRoute
  ServerPortIndexRoute: typeof ServerPortIndexRoute
  TestsIndexRoute: typeof TestsIndexRoute
  ServerGroupsManageServersGroupIdRoute: typeof ServerGroupsManageServersGroupIdRoute
  ConfigNatsIndexRoute: typeof ConfigNatsIndexRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexRoute: IndexRoute,
  AboutRoute: AboutRoute,
  ServerGroupsAddRoute: ServerGroupsAddRoute,
  ServerGroupsAddServerRoute: ServerGroupsAddServerRoute,
  ServerGroupsIncidentsRoute: ServerGroupsIncidentsRoute,
  ServerGroupsOldindexRoute: ServerGroupsOldindexRoute,
  ServerPortFormRoute: ServerPortFormRoute,
  ServerPortStatusRoute: ServerPortStatusRoute,
  ServerGroupsIndexRoute: ServerGroupsIndexRoute,
  ServerPortIndexRoute: ServerPortIndexRoute,
  TestsIndexRoute: TestsIndexRoute,
  ServerGroupsManageServersGroupIdRoute: ServerGroupsManageServersGroupIdRoute,
  ConfigNatsIndexRoute: ConfigNatsIndexRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/about",
        "/server-groups/add",
        "/server-groups/add-server",
        "/server-groups/incidents",
        "/server-groups/oldindex",
        "/server-port/form",
        "/server-port/status",
        "/server-groups/",
        "/server-port/",
        "/tests/",
        "/server-groups/manage-servers/$groupId",
        "/config/nats/"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/about": {
      "filePath": "about.tsx"
    },
    "/server-groups/add": {
      "filePath": "server-groups/add.tsx"
    },
    "/server-groups/add-server": {
      "filePath": "server-groups/add-server.tsx"
    },
    "/server-groups/incidents": {
      "filePath": "server-groups/incidents.tsx"
    },
    "/server-groups/oldindex": {
      "filePath": "server-groups/oldindex.tsx"
    },
    "/server-port/form": {
      "filePath": "server-port/form.tsx"
    },
    "/server-port/status": {
      "filePath": "server-port/status.tsx"
    },
    "/server-groups/": {
      "filePath": "server-groups/index.tsx"
    },
    "/server-port/": {
      "filePath": "server-port/index.tsx"
    },
    "/tests/": {
      "filePath": "tests/index.tsx"
    },
    "/server-groups/manage-servers/$groupId": {
      "filePath": "server-groups/manage-servers.$groupId.tsx"
    },
    "/config/nats/": {
      "filePath": "config/nats/index.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
