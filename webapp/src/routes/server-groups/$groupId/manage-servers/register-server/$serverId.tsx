import { createFileRoute, Link } from '@tanstack/react-router'
import { useState, useEffect } from 'react'
import { Server, ServerRegStatusEnum } from '@/components/types/ServerGroup'
import { Button } from '@/components/ui/button'
import { ClipboardCopy } from 'lucide-react'

export const Route = createFileRoute(
  '/server-groups/$groupId/manage-servers/register-server/$serverId',
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
    ip: '',
    agent_port: -1,
    agent_version: '',
    fqdn: '',
    reg_status: ServerRegStatusEnum.New,
    desc: '',
  })

  useEffect(() => {
    ; (async () => {
      const serverResp = await fetch(
        `/api/v1/server-groups/${groupId}/servers/${serverId}`,
      )
      const serverData = await serverResp.json()
      console.table(serverData)
      setServer(serverData)
    })()
  }, [serverId])

  useEffect(() => {
    if (!showToken) return

    let interval_id = setInterval(() => {
      ; (async () => {
        const serverResp = await fetch(
          `/api/v1/server-groups/${groupId}/servers/${serverId}`,
        )
        const serverData = await serverResp.json()
        if (serverData.reg_status === ServerRegStatusEnum.Active) {
          setServer({ ...serverData, reg_status: ServerRegStatusEnum.Active })
          clearInterval(interval_id)
        }
      })()
    }, 1000);

  }, [showToken])

  return (
    <>
      <div className="w-2/3 max-w-xl mb-8 mx-8 bg-muted/50 rounded-xl p-8">
        <h2 className="pb-4 text-3xl font-semibold">
          Register Server: {server.ip}
        </h2>
        <h4 className="pb-2 text-lg font-medium">
          This is a one-time setup process to register the server with the
          agent.
        </h4>
      </div>
      <Button
        className="w-36 bg-blue-500 hover:bg-blue-700 text-white ml-8"
        size={'sm'}
        variant={'outline'}
        onClick={async () => {
          const tokenResp = await fetch(`/api/v1/server-groups/${groupId}/servers/${serverId}/regtoken`)
          const _tokenData = await tokenResp.json()
          const tokenData = _tokenData.token
          const tokenURL =
            window.location.origin +
            `/api/v1/server/${serverId}/verifytoken/${tokenData}`
          console.log(tokenURL)
          setToken(tokenURL)
          setShowToken(true)
        }}
      >
        Generate Token
      </Button>
      {showToken && (
        <div className="m-8 bg-muted/50 rounded-xl p-8">
          <h4 className="pb-2 text-lg font-medium">
            Registration URL: {token} &nbsp;&nbsp;
            <ClipboardCopy
              className='inline-block w-8 h-8 p-1 bg-gray-500 text-white cursor-pointer hover:bg-gray-600'
              onClick={() => navigator.clipboard.writeText(token)}
              role="button"
            />
          </h4>
          <h6>Enter the entire URL on the client</h6>
        </div>
      )}
    </>
  )
}
