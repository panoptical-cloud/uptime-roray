import { ColumnDef } from "@tanstack/react-table"

export type Server = {
    name: string
    hostname: number
    ip: string
    monitor: boolean
    clientVersion: string
}

export const columns: ColumnDef<Server>[] = [
    {
        accessorKey: "status",
        header: "Status",
    },
    {
        accessorKey: "email",
        header: "Email",
    },
    {
        accessorKey: "amount",
        header: "Amount",
    },
]
