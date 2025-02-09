import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/servers/incidents')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/servers/incidents"!</div>
}
