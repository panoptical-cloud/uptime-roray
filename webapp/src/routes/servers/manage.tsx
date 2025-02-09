import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/servers/manage')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/servers/manage"!</div>
}
