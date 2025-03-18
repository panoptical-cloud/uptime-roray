import { useEffect, useState } from 'react'
import { createFileRoute, Link } from '@tanstack/react-router'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Progress } from '@/components/ui/progress'
import { Badge } from '@/components/ui/badge'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet'
import { Server, Cloud } from 'lucide-react'
import { Separator } from '@/components/ui/separator'
import { ScrollArea } from '@/components/ui/scroll-area'

export const Route = createFileRoute('/server-groups/')({
  component: RouteComponent,
})

function RouteComponent() {
  const [progress, setProgress] = useState(13)
  const [servers, setServers] = useState<Server[]>([])


  useEffect(() => {
    const timer = setTimeout(() => setProgress(66), 500)
    return () => clearTimeout(timer)
  }, [])

  const serverData = [
    {
      name: 'FE-NGINX-1',
      host: 'abc.xyz.com',
      ip: '12.232.54.12',
      status: {
        online: true,
        since: '2 days 3 hrs',
      },
      cpu: '45% used',
      ram: '6.2/16 GB',
      disk: '234/500 GB',
    },
    // Add more server entries as needed
  ]

  const [serversa, setServersa] = useState(serverData)

  return (
    <>
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        <div className="grid auto-rows-min gap-4 md:grid-cols-3">
          <div className="aspect-video rounded-xl bg-muted/50">
            <h2 className="scroll-m-20 text-3xl font-semibold tracking-tight p-4">
              Overview
            </h2>
          </div>
          <div className="aspect-video rounded-xl bg-muted/50">
            <h4 className="scroll-m-20 text-xl font-semibold tracking-tight p-4">
              Events
            </h4>
          </div>
          <div className="aspect-video rounded-xl bg-muted/50">
            <h4 className="scroll-m-20 text-xl font-semibold tracking-tight p-4">
              Incidents
            </h4>
          </div>
        </div>
        <div className="min-h-[100vh] flex-1 rounded-xl bg-muted/50 md:min-h-min">
          <div className="p-6">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Name</TableHead>
                  <TableHead>Host</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>CPU Usage</TableHead>
                  <TableHead>RAM Usage</TableHead>
                  <TableHead>Disk Usage</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {serversa.map((server) => (
                  <TableRow key={server.name}>
                    <TableCell>
                      <span className="font-semibold"> {server.name}</span>
                      <Sheet>
                        <SheetTrigger asChild>
                          <p className="text-xs pt-1 text-blue-500 underline underline-offset-4">
                            <Button variant="ghost" size="icon">
                              Details
                            </Button>
                          </p>
                        </SheetTrigger>

                        <SheetContent>
                          <SheetHeader>
                            <SheetTitle>Server: FE-NGINX-1</SheetTitle>
                            <SheetDescription>
                              Agent: v 0.12 [online]
                            </SheetDescription>
                            <h3>Last 24 hours</h3>
                          </SheetHeader>
                          <Separator className="mb-4 mt-4" />
                          <div className="grid gap-4 mb-4">
                            <h4>
                              <Badge variant="secondary" className="text-sm">
                                2 Events
                              </Badge>
                            </h4>
                            <ScrollArea className="h-[260px] w-[340px] rounded-md border p-4 bg-muted">
                              <div className="mt-2">
                                {/* <h4 className="text-lg font-semibold mb-2">Timeline</h4> */}
                                <ul className="relative border-l border-gray-200 ml-4">
                                  <li className="mb-4 -ml-4 flex items-start">
                                    <div className="mr-2">
                                      <Server className="w-8 h-8 p-1 bg-orange-500 text-white rounded-md" />
                                    </div>
                                    <div>
                                      <time className="mb-1 text-sm font-normal text-gray-400">
                                        09:00 AM [2H 31M ago]
                                      </time>
                                      <h3 className="text-base font-semibold text-gray-900">
                                        Server Booted
                                      </h3>
                                      <p className="mb-4 text-base font-normal text-gray-500">
                                        The server started successfully.
                                      </p>
                                    </div>
                                  </li>
                                  <li className="mb-2 -ml-4 flex items-start">
                                    <div className="mr-2">
                                      <Cloud className="w-8 h-8 p-1 bg-gray-500 text-white rounded-md" />
                                    </div>
                                    <div>
                                      <time className="mb-1 text-sm font-normal text-gray-400">
                                        12:30 PM
                                      </time>
                                      <h3 className="text-base font-semibold text-gray-900">
                                        Deployment Completed
                                      </h3>
                                      <p className="mb-4 text-base font-normal text-gray-500">
                                        New changes were deployed.
                                      </p>
                                    </div>
                                  </li>
                                </ul>
                              </div>
                            </ScrollArea>
                            <Separator />
                            <h4>
                              <Badge variant="secondary" className="text-sm">
                                3 Incidents
                              </Badge>
                            </h4>
                            <ScrollArea className="h-[260px] w-[340px] rounded-md border p-4 bg-muted">
                              <div className="mt-2">
                                {/* <h4 className="text-lg font-semibold mb-2">Timeline</h4> */}
                                <ul className="relative border-l border-gray-200 ml-4">
                                  <li className="mb-4 -ml-4 flex items-start">
                                    <div className="mr-2">
                                      <Server className="w-8 h-8 p-1 bg-orange-500 text-white rounded-md" />
                                    </div>
                                    <div>
                                      <time className="mb-1 text-sm font-normal text-gray-400">
                                        09:00 AM [2H 31M ago]
                                      </time>
                                      <h3 className="text-base font-semibold text-gray-900">
                                        Server Booted
                                      </h3>
                                      <p className="mb-4 text-base font-normal text-gray-500">
                                        The server started successfully.
                                      </p>
                                    </div>
                                  </li>
                                  <li className="mb-4 -ml-4 flex items-start">
                                    <div className="mr-2">
                                      <Cloud className="w-8 h-8 p-1 bg-gray-500 text-white rounded-md" />
                                    </div>
                                    <div>
                                      <time className="mb-1 text-sm font-normal text-gray-400">
                                        12:30 PM
                                      </time>
                                      <h3 className="text-base font-semibold text-gray-900">
                                        Deployment Completed
                                      </h3>
                                      <p className="mb-4 text-base font-normal text-gray-500">
                                        New changes were deployed.
                                      </p>
                                    </div>
                                  </li>
                                  <li className="mb-4 -ml-4 flex items-start">
                                    <div className="mr-2">
                                      <Cloud className="w-8 h-8 p-1 bg-gray-500 text-white rounded-md" />
                                    </div>
                                    <div>
                                      <time className="mb-1 text-sm font-normal text-gray-400">
                                        12:30 PM
                                      </time>
                                      <h3 className="text-base font-semibold text-gray-900">
                                        Deployment Completed
                                      </h3>
                                      <p className="mb-4 text-base font-normal text-gray-500">
                                        New changes were deployed.
                                      </p>
                                    </div>
                                  </li>
                                </ul>
                              </div>
                            </ScrollArea>
                          </div>
                          <SheetFooter>
                            <SheetClose>
                              <Button type="submit">Close</Button>
                              <Button type="submit">Close</Button>
                            </SheetClose>
                          </SheetFooter>
                        </SheetContent>
                      </Sheet>
                    </TableCell>
                    <TableCell>
                      {server.host}
                      <p className="text-xs text-muted-foreground pt-1">
                        IP: 12.23.43.11
                      </p>
                    </TableCell>
                    <TableCell>
                      <Badge
                        variant={
                          server.status.online ? 'default' : 'destructive'
                        }
                        color="bg-primary"
                        className="bg-green-400"
                      >
                        {server.status.online ? 'Online' : 'Offline'}
                      </Badge>
                      <p className="text-xs text-muted-foreground pt-1">
                        Since: 2D 13H 43M 21S
                      </p>
                    </TableCell>
                    <TableCell>
                      {server.cpu}
                      <p>
                        <Progress
                          value={progress}
                          className="w-[60%] h-2 bg-slate-300"
                        />
                      </p>
                      <p className="text-xs text-muted-foreground pt-1">
                        6 Core Intel Xeon 2.4 GHz
                      </p>
                    </TableCell>
                    <TableCell>
                      63%
                      <p>
                        <Progress
                          value={progress}
                          className="w-[60%] h-2 bg-slate-300"
                        />
                      </p>
                      <p className="text-xs text-muted-foreground pt-1">
                        {server.ram}
                      </p>
                    </TableCell>
                    <TableCell>
                      72%
                      <p>
                        <Progress
                          value={progress}
                          className="w-[60%] h-2 bg-slate-300"
                        />
                      </p>
                      <p className="text-xs text-muted-foreground pt-1">
                        {server.disk}
                      </p>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </div>
      </div>
    </>
  )
}
