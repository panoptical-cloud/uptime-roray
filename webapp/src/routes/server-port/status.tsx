import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/server-port/status')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/server-port/status"!</div>
}
