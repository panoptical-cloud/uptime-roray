import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/server-port/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/server-port/"!</div>
}
