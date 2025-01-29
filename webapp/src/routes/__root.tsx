import * as React from 'react'
import { Link, Outlet, createRootRoute } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import "./index.css"
export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {
  return (
    <>
      <div className="p-2 flex gap-2 text-lg">
        <Link
          to="/"
          activeProps={{
            className: 'font-bold',
          }}
          activeOptions={{ exact: true }}
        >
          Dashboard
        </Link>{' '}
        <Link
          to="/servers"
          activeProps={{
            className: 'font-bold',
          }}
        >
          Servers
        </Link>
        <Link
          to="/ports"
          activeProps={{ className: 'font-bold' }}
        >
          Ports
        </Link>
        <Link
          to="/urls"
          activeProps={{ className: 'font-bold' }}
        >
          URLs
        </Link>
      </div>
      <hr />
      <Outlet />
      <TanStackRouterDevtools position="bottom-right" />
    </>
  )
}
