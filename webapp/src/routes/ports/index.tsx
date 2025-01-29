import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/ports/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/ports/"!</div>
}
