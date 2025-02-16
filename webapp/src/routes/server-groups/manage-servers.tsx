import { Button } from '@/components/ui/button'
import { createFileRoute, Link } from '@tanstack/react-router'

export const Route = createFileRoute('/server-groups/manage-servers')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="w-2/3 max-w-xl m-8 bg-muted/50 rounded-xl p-8">
      <h2 className="pb-4 text-3xl font-semibold">Group: QA-App1-Svc1</h2>
      <h4 className="pb-2 text-lg font-medium">Description: Desc of QA-App1-Svc1</h4>
      <div className="flex justify-end ">
        <Link to="/server-groups/add-server" >
          <Button className="w-36 bg-blue-500 hover:bg-blue-700 text-white" size={'sm'} variant={'secondary'} >
            Add New Server
          </Button>
        </Link>
      </div>
    </div>
  )
}
