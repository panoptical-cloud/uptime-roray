import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { createFileRoute, Link } from '@tanstack/react-router'
import { useEffect, useState } from 'react'

import type { Server, ServerGroup } from '@/components/types/ServerGroup'

export const Route = createFileRoute('/server-groups/manage-servers/$groupId/')(
  {
    component: RouteComponent,
  },
)

function RouteComponent() {
  const { groupId } = Route.useParams()
  
  const [servers, setServers] = useState<Server[]>([])
  const [serverGroup, setServerGroup] = useState<ServerGroup>({
    id: -1,
    name: '',
    desc: '',
  })

  useEffect(() => {
    (async () => {
      const serverGroup = await fetch(`/api/v1/server-groups/${groupId}`)
      const serverGroupData = await serverGroup.json()
      console.table(serverGroupData)
      setServerGroup(serverGroupData)
      const serversResp = await fetch(
        `/api/v1/server-groups/${groupId}/servers`,
      )
      const serversList = await serversResp.json()
      console.table(serversList)
      setServers(serversList)
    })()
  }, [groupId])

  return (
    <>
      <div className="w-2/3 max-w-xl mb-8 mx-8 bg-muted/50 rounded-xl p-8">
        <h2 className="pb-4 text-3xl font-semibold">
          Group: {serverGroup.name}
        </h2>
        <h4 className="pb-2 text-lg font-medium">
          Description: {serverGroup.desc}
        </h4>
        <div className="flex justify-end ">
          <Link to="/server-groups/add-server">
            <Button
              className="w-36 bg-blue-500 hover:bg-blue-700 text-white"
              size={'sm'}
              variant={'secondary'}
            >
              Add New Server
            </Button>
          </Link>
        </div>
      </div>

      <div className=" flex-1 rounded-xl bg-muted/50 max-w-7xl mx-8 p-8">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Hostname</TableHead>
              <TableHead>IP</TableHead>
              <TableHead>Agent Port</TableHead>
              <TableHead>Agent Version</TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {servers.map((server) => (
              <TableRow key={server.id}>
                <TableCell className="py-2 px-4">{server.name}</TableCell>
                <TableCell className="py-2 px-4">{server.hostname}</TableCell>
                <TableCell className="py-2 px-4">{server.ip}</TableCell>
                <TableCell className="py-2 px-4">{server.agent_port}</TableCell>
                <TableCell className="py-2 px-4">
                  {server.agent_version}
                </TableCell>
                <TableCell className="py-2 px-4">
                  <DropdownMenu>
                    <DropdownMenuTrigger>
                      <Button size="sm" variant="ghost">
                        ...
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                      <DropdownMenuItem asChild>
                        <Link
                          to={`/server-groups/manage-servers/${serverGroup.id}/register-server/${server.id}`}
                        >
                          Register
                        </Link>
                      </DropdownMenuItem>
                      <DropdownMenuItem asChild>
                        <Link
                          to={`/server-groups/manage-servers/${serverGroup.id}/edit-server/${server.id}`}
                        >
                          Edit
                        </Link>
                      </DropdownMenuItem>
                      <DropdownMenuItem asChild>
                        <Link
                          to={`/server-groups/manage-servers/${serverGroup.id}/delete-server/${server.id}`}
                        >
                          Delete
                        </Link>
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </>
  )
}
