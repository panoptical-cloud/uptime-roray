import { createFileRoute, Link } from '@tanstack/react-router'
import { useState, useEffect } from 'react'
import type { Server } from '@/components/types/ServerGroup'
import { Button } from '@/components/ui/button'

export const Route = createFileRoute(
    '/server-groups/manage-servers/$groupId/register-server/$serverId',
)({
    component: RouteComponent,
})

function RouteComponent() {
    const { groupId, serverId } = Route.useParams()
    const [token, setToken] = useState<string>('')
    const [showToken, setShowToken] = useState<boolean>(false)

    const [server, setServer] = useState<Server>({
        id: '',
        name: '',
        hostname: '',
        ip: '',
        agent_port: -1,
        agent_version: '',
    })

    useEffect(() => {
        (async () => {
            const serverResp = await fetch(`/api/v1/server-groups/${groupId}/servers/${serverId}`)
            const serverData = await serverResp.json()
            console.table(serverData)
            setServer(serverData)
        }
        )()
    }, [serverId])

    return (
        <>
            <div className="w-2/3 max-w-xl mb-8 mx-8 bg-muted/50 rounded-xl p-8">
                <h2 className="pb-4 text-3xl font-semibold">
                    Register Server: {server.hostname}
                </h2>
                <h4 className="pb-2 text-lg font-medium">
                    This is a one-time setup process to register the server with the agent.
                </h4>
            </div>
            <Button
                className="w-36 bg-blue-500 hover:bg-blue-700 text-white ml-8"
                size={'sm'}
                variant={'outline'}
                onClick={async () => {
                    const tokenResp = await fetch(`/api/v1/server/${serverId}/regtoken`)
                    const _tokenData = await tokenResp.json()
                    const tokenData = _tokenData.token
                    const tokenURL = window.location.origin + `/api/v1/server/${serverId}/verifytoken/${tokenData}`
                    console.log(tokenURL)
                    setToken(tokenURL)
                    setShowToken(true)
                }}
            >
                Generate Token
            </Button>
            {
                showToken && (
                    <div className="m-8 bg-muted/50 rounded-xl p-8">
                        <h4 className="pb-2 text-lg font-medium">
                            Registration URL: {token}
                        </h4>
                        <h6>
                            Enter the entire URL on the client
                        </h6>
                    </div>
                )
            }
        </>
    )
}
