import { createFileRoute } from '@tanstack/react-router'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Badge } from '@/components/ui/badge'
import { useState } from 'react'

export const Route = createFileRoute('/servers/overview')({
    component: RouteComponent,
})

function RouteComponent() {
    const serverData = [
        {
            name: "Server-01",
            cpu: "45%",
            ram: "6.2GB/16GB",
            disk: "234GB/500GB",
            uptime: "15d 4h",
            status: "healthy"
        },
        // Add more server entries as needed
    ]

    const [servers, setServers] = useState(serverData)

    return (
        <div className="p-6">
            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead>Server Name</TableHead>
                        <TableHead>CPU Usage</TableHead>
                        <TableHead>RAM Usage</TableHead>
                        <TableHead>Disk Space</TableHead>
                        <TableHead>Uptime</TableHead>
                        <TableHead>Status</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {servers.map((server) => (
                        <TableRow key={server.name}>
                            <TableCell>{server.name}</TableCell>
                            <TableCell>{server.cpu}</TableCell>
                            <TableCell>{server.ram}</TableCell>
                            <TableCell>{server.disk}</TableCell>
                            <TableCell>{server.uptime}</TableCell>
                            <TableCell>
                                <Badge
                                    variant={server.status === 'healthy' ? 'default' : 'destructive'}
                                    color='green'
                                >
                                    {server.status}
                                </Badge>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </div>
    )
}
