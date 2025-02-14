import { Button } from '@/components/ui/button'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/servers/manage/')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="w-5/6 m-5 p-4">
      <h2 className="pb-2 text-3xl font-semibold">
        Manage Servers
      </h2>
      <div className="flex justify-end ">
        <Button className="w-28" size={'sm'} variant={'secondary'}>
          Add Server
        </Button>
      </div>
    </div>
  )
}
