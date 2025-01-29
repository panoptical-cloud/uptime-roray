import { createFileRoute, Link } from '@tanstack/react-router'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

import { ChevronUp } from "lucide-react";
import { Button } from '@/components/ui/button';
import { Separator } from "@/components/ui/separator";

export const Route = createFileRoute('/servers/')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <>
      <h4>Servers Index</h4>
      <Button variant="default" asChild>
        <Link to="/servers/new">Add New Server</Link>
      </Button>
      <Card className="w-[350px]">
        <CardHeader>
          <CardTitle>Roray Blog Server</CardTitle>
          <CardDescription>
            <span style={{ display: 'flex', alignItems: 'center' }}><ChevronUp color='lightgreen' size="24" /> Healthy</span>
          </CardDescription>
        </CardHeader>
        <CardContent>
          <p>Utilization</p>
          <div className="flex-1 space-y-1">
            <p>Disk: 48%</p>
            <p>RAM: 23%</p>
            <p>CPU: 72%</p>
          </div>
        </CardContent>
        <CardFooter>
          <div className="flex justify-between">
            <Button variant="secondary" asChild >
              <Link to="/servers/1">Details</Link>
            </Button>
            <Button variant="link" asChild >
              <Link to="/servers/1/edit">Edit</Link>
            </Button>
            <Button variant="destructive" asChild >
              <Link to="/servers/1/delete">Delete</Link>
            </Button>
          </div>
        </CardFooter>
      </Card>
    </>
  )

}
