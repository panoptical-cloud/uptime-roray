import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute(
  '/server-groups/manage-servers/$groupId/register-server/$serverId',
)({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div>
      Hello "/server-groups/manage-servers/$groupId/register-server/$serverId"!
    </div>
  )
}
