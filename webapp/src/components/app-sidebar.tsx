import * as React from "react"
import {
  AudioWaveform,
  BookOpen,
  Database,
  Command,
  Frame,
  GalleryVerticalEnd,
  Map,
  PieChart,
  Settings2,
  Server,
} from "lucide-react"
import { NavServerGroups } from "@/components/nav-server-groups"
import { NavProjects } from "@/components/nav-projects"
import { NavUser } from "@/components/nav-user"
import { TeamSwitcher } from "@/components/team-switcher"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar"

// This is sample data.
const data = {
  user: {
    name: "shadcn",
    email: "m@example.com",
    avatar: "/avatars/shadcn.jpg",
  },
  teams: [
    {
      name: "Acme Inc",
      logo: GalleryVerticalEnd,
      plan: "Enterprise",
    },
    {
      name: "Acme Corp.",
      logo: AudioWaveform,
      plan: "Startup",
    },
    {
      name: "Myra Corp.",
      logo: Command,
      plan: "Free",
    },
  ],
  // navServers: [
  //   {
  //     title: "QA-App1-Svc1",
  //     url: "#",
  //     icon: Server,
  //     isActive: true,
  //     items: [
  //       {
  //         title: "Overview",
  //         url: "/server-groups",
  //       },
  //       {
  //         title: "Incidents",
  //         url: "#",
  //       },
  //       {
  //         title: "Manage Servers",
  //         url: "/server-groups/manage-servers",
  //       },
  //     ],
  //   },
  //   {
  //     title: "Prod-App1-Svc1",
  //     url: "#",
  //     icon: Server,
  //     items: [
  //       {
  //         title: "Overview",
  //         url: "/servers/overview",
  //       },
  //       {
  //         title: "Incidents",
  //         url: "#",
  //       },
  //       {
  //         title: "Manage Servers",
  //         url: "#",
  //       },
  //     ],
  //   },
  //   {
  //     title: "Documentation",
  //     url: "#",
  //     icon: BookOpen,
  //     items: [
  //       {
  //         title: "Introduction",
  //         url: "#",
  //       },
  //       {
  //         title: "Get Started",
  //         url: "#",
  //       },
  //       {
  //         title: "Tutorials",
  //         url: "#",
  //       },
  //       {
  //         title: "Changelog",
  //         url: "#",
  //       },
  //     ],
  //   },
  //   {
  //     title: "Settings",
  //     url: "#",
  //     icon: Settings2,
  //     items: [
  //       {
  //         title: "General",
  //         url: "#",
  //       },
  //       {
  //         title: "Team",
  //         url: "#",
  //       },
  //       {
  //         title: "Billing",
  //         url: "#",
  //       },
  //       {
  //         title: "Limits",
  //         url: "#",
  //       },
  //     ],
  //   },
  // ],
  projects: [
    {
      name: "Design Engineering",
      url: "#",
      icon: Frame,
    },
    {
      name: "Sales & Marketing",
      url: "#",
      icon: PieChart,
    },
    {
      name: "Travel",
      url: "#",
      icon: Map,
    },
  ],
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const [navData, setNavData] = React.useState([])

  React.useEffect(() => {
    (async () => {
      const _resp = await fetch('/api/v1/server-groups')
      const qr = await _resp.json()
      qr.map((k, i) => {
        qr[i].url = "#"
        qr[i].icon = Server
        qr[i].items = []
        qr[i].items.push({
          title: "Overview",
          url: `/server-groups/${k.id}`,
        }, {
          title: "Incidents",
          url: "#",
        }, {
          title: "Manage Servers",
          url: `/server-groups/${k.id}/manage-servers`,
        })
      })
      // console.table(qr)
      setNavData(qr)
    })()
  }, [])

  return (
    <Sidebar collapsible="icon" {...props} variant="floating">
      <SidebarHeader>
        <TeamSwitcher teams={data.teams} />
      </SidebarHeader>
      <SidebarContent>
        {/* {typeof navData != "undefined" && navData !== null && <NavServerGroups items={navData} />} */}
        {typeof navData != "undefined" && navData !== null && <NavServerGroups items={navData} />}
        <NavProjects projects={data.projects} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}
