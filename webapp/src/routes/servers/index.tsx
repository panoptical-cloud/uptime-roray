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
import { useEffect, useState } from 'react';

export const Route = createFileRoute('/servers/')({
  component: RouteComponent,
})

function RouteComponent() {
  const [serverStatus, setServerStatus] = useState({ "disk": 0, "ram": 0, "cpu": 0 })

  useEffect(() => {
    console.log("Component useEffect mounted");

    // opening a connection to the server to begin receiving events from it
    const eventSource = new EventSource("http://localhost:9191/api/v1/events/server-status");

    eventSource.addEventListener("server-1", (event) => {
      if (event.data) {
        const serverStatusData = JSON.parse(event.data);
        setServerStatus({ ...serverStatusData });
      }
    });

    // attaching a handler to receive message events
    // eventSource.onmessage = (event) => {
    //   console.log("Received event: ", event.data);
    //   const serverStatusData = JSON.parse(event.data);
    //   console.log(serverStatusData);
    //   setServerStatus({ ...serverStatusData });
    // };

    // attaching a handler to handle errors
    eventSource.onerror = (error) => {
      console.error("EventSource failed: ", error);
      return () => {
        console.log("ES closing since in error");
        eventSource.close();
      };
    };

    // terminating the connection on component unmount
    return () => {
      console.log("Component unmounted, closing EventSource");
      eventSource.close();
    };
  }, []);

  return (
    <>
      <Button variant="default" asChild>
        {/* <Link to="/servers/new">Add New Server</Link> */}
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
            <p>Disk: {serverStatus.disk} %</p>
            <p>RAM: {serverStatus.ram} %</p>
            <p>CPU: {serverStatus.cpu} %</p>
          </div>
        </CardContent>
        <CardFooter>
          <div className="flex justify-between">
            <Button variant="secondary" asChild >
              {/* <Link to="/servers/1">Details</Link> */}
            </Button>
            <Button variant="link" asChild >
              {/* <Link to="/servers/1/edit">Edit</Link> */}
            </Button>
            <Button variant="destructive" asChild >
              {/* <Link to="/servers/1/delete">Delete</Link> */}
            </Button>
          </div>
        </CardFooter>
      </Card>
    </>
  )

}
