import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/urls/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/urls/"!</div>
}
