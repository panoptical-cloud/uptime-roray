import type { FieldApi } from '@tanstack/react-form'
import { useForm } from '@tanstack/react-form'
import { createFileRoute } from '@tanstack/react-router'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { useMutation } from "react-query";
import FieldInfo from '@/components/FieldInfo';

export const Route = createFileRoute('/server-groups/add')({
  component: RouteComponent,
})

function RouteComponent() {
  // Initialize the mutation
  const mutation = useMutation(
    async (formData: { groupName: string; groupDesc: string }) => {
      const response = await fetch('/api/v1/server-groups', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: formData.groupName,
          desc: formData.groupDesc,
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
      },
    }
  )

  const form = useForm({
    defaultValues: {
      groupName: '',
      groupDesc: '',
    },
    onSubmit: async ({ value }) => {
      // Trigger the mutation on form submit
      mutation.mutate({
        groupName: value.groupName,
        groupDesc: value.groupDesc
      }, {
        onSuccess: () => {
          form.reset()
        }
      })
    }
  })

  return (
    <div className="w-2/3 max-w-xl m-8 bg-muted/50 rounded-xl p-8">
      <h2 className="mb-8 text-4xl font-semibold ">Add new server group</h2>
      <form
        onSubmit={(e) => {
          e.preventDefault()
          e.stopPropagation()
          form.handleSubmit()
        }}
      >
        <div>
          <form.Field
            name="groupName"
            validators={{
              onChangeAsync: ({ value }) =>
                !value
                  ? 'A group name is required'
                  : value.length < 3
                    ? 'group name must be at least 3 characters'
                    : undefined,
              onChangeAsyncDebounceMs: 500,
            }}
            children={(field) => (
              <div className="flex flex-col space-y-4 m-4">
                <Label htmlFor={field.name}>Group Name</Label>
                <Input
                  className="pb-4"
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                  type="text"
                />
                <FieldInfo field={field} />
              </div>
            )}
          />
        </div>
        <div>
          <form.Field
            name="groupDesc"
            validators={{
              onChangeAsync: ({ value }) =>
                !value
                  ? 'A group description is required'
                  : value.length < 3
                    ? 'group description must be at least 3 characters'
                    : undefined,
              onChangeAsyncDebounceMs: 500,
            }}
            children={(field) => (
              <div className="flex flex-col space-y-4 m-4">
                <Label htmlFor={field.name}>Group Description</Label>
                <Input
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                  type="text"
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
                disabled={!canSubmit || isSubmitting}
                className="w-28 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mx-4"
              >
                {isSubmitting ? 'Submitting...' : 'Submit'}
              </button>
              <button type="reset" onClick={() => form.reset()}>
                Reset
              </button>
            </div>
          )}
        />
      </form>
    </div>
  )
}