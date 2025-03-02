import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { useForm, type FieldApi } from '@tanstack/react-form'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { useMutation } from 'react-query'
import FieldInfo from '@/components/FieldInfo'

export const Route = createFileRoute(
  '/server-groups/$groupId/manage-servers/add-server',
)({
  component: RouteComponent,
})

function RouteComponent() {
  const { groupId } = Route.useParams()

  const navigate = useNavigate();
  const mutation = useMutation(
    async (formData: {
      name: string
      ip: string
      group_id: number
    }) => {
      const response = await fetch('/api/v1/server-groups/servers', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: formData.name,
          ip: formData.ip,
          group_id: formData.group_id,
        }),
      })
      return response.json()
    },
    {
      onError: (err) => {
        console.log(err)
      },
      onSuccess: (data) => {
        console.log(data)
        navigate({ to: `/server-groups/${groupId}/manage-servers` });
      },
    },
  )

  const form = useForm({
    defaultValues: {
      name: '',
      ip: '',
    },
    onSubmit: async ({ value }) => {
      // Trigger the mutation on form submit
      mutation.mutate(
        {
          name: value.name,
          ip: value.ip,
          group_id: parseInt(groupId),
        },
        {
          onSuccess: () => {
            form.reset()
          },
          onError: (err) => {
            console.log(err)
          },
        },
      )
    },
  })

  return (
    <div className="w-2/3 max-w-xl m-8 bg-muted/50 rounded-xl p-8">
      <h2 className="mb-8 text-4xl font-semibold ">
        Add new server to 'group_name'
      </h2>
      <form
        onSubmit={(e) => {
          e.preventDefault()
          e.stopPropagation()
          form.handleSubmit()
        }}
      >
        <div>
          <form.Field
            name="name"
            validators={{
              onChangeAsync: ({ value }) =>
                !value
                  ? 'A server name is required'
                  : value.length < 3
                    ? 'server name must be at least 3 characters'
                    : undefined,
              onChangeAsyncDebounceMs: 300,
            }}
            children={(field) => (
              <div className="flex flex-col space-y-4 m-4">
                <Label htmlFor={field.name}>Friendly Name</Label>
                <Input
                  className="pb-4"
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                />
                <FieldInfo field={field} />
              </div>
            )}
          />
        </div>
        <div>
          <form.Field
            name="ip"
            validators={{
              onChangeAsync: ({ value }) =>
                !value
                  ? 'An IP address is required'
                  : value.length < 3
                    ? 'IP address must be at least 3 characters'
                    : undefined,
              onChangeAsyncDebounceMs: 300,
            }}
            children={(field) => (
              <div className="flex flex-col space-y-4 m-4">
                <Label htmlFor={field.name}>IP Address</Label>
                <Input
                  className="pb-4"
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                />
                <FieldInfo field={field} />
              </div>
            )}
          />
        </div>
        <form.Subscribe
          selector={(state) => [state.canSubmit, state.isSubmitting]}
          children={([canSubmit, isSubmitting]) => (
            <div className="flex justify-end">
              <button
                type="submit"
                className="w-28 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mx-4"
                disabled={!canSubmit || isSubmitting}
              >
                {isSubmitting ? 'Adding...' : 'Add'}
              </button>
            </div>
          )}
        />
      </form >
    </div >
  )
}
