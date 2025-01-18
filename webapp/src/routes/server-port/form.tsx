import { createFileRoute } from '@tanstack/react-router'
import type { FieldApi } from '@tanstack/react-form'
import { useForm } from '@tanstack/react-form'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

function FieldInfo({ field }: { field: FieldApi<any, any, any, any> }) {
  return (
    <>
      {field.state.meta.isTouched && field.state.meta.errors.length ? (
        <em>{field.state.meta.errors.join(',')}</em>
      ) : null}
      {field.state.meta.isValidating ? 'Validating...' : null}
    </>
  )
}

export const Route = createFileRoute('/server-port/form')({
  component: RouteComponent,
})

function RouteComponent() {
  const form = useForm({
    defaultValues: {
      server: '',
      port: -1,
    },
    onSubmit: async ({ value }) => {
      // Do something with form data
      console.log(value)
    },
  })

  return (
    <div>
      <h2>
        Add a server port for monitoring
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
            name="server"
            validators={{
              onChangeAsync: ({ value }) =>
                !value
                  ? 'A server name is required'
                  : value.length < 3
                    ? 'server name must be at least 3 characters'
                    : undefined,
              onChangeAsyncDebounceMs: 500,
            }}
            children={(field) => (
              <>
                <Label htmlFor={field.name}>Server Host/IP</Label>
                <Input
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onChange={e => field.handleChange(e.target.value)}
                  type="text" />
                <FieldInfo field={field} />
              </>
            )}
          />
        </div>
        <div>
          <form.Field
            name="port"
            validators={{
              onChangeAsync: ({ value }) => value < 0 ? 'Port must be a positive number' : value > 65535 ? 'Port must be less than 65536' : undefined,
              onChangeAsyncDebounceMs: 500,
              // onChangeAsync: async ({ value }) => {
              //   await new Promise(resolve => setTimeout(resolve, 1000))
              //   return "Enter valid port number"
              // }
            }}
            children={(field) => (
              <>
                <Label htmlFor={field.name}>Port</Label>
                <Input
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onChange={e => field.handleChange(Number(e.target.value))}
                  type="number" />
                <FieldInfo field={field} />
              </>
            )}
          />
        </div>
        <form.Subscribe
          selector={(state) => [state.canSubmit, state.isSubmitting]}
          children={([canSubmit, isSubmitting]) => (
            <>
              <button type="submit" disabled={!canSubmit}>
                {isSubmitting ? '...' : 'Submit'}
              </button>
              <button type="reset" onClick={() => form.reset()}>
                Reset
              </button>
            </>
          )}
        />
      </form>
    </div>
  )
}
