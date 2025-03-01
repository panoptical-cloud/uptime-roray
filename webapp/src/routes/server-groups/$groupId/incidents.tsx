import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/server-groups/$groupId/incidents')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/servers/incidents"!</div>
}
