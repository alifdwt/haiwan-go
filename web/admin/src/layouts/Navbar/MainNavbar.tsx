import { cn } from "@/lib/utils";

export default function MainNav({
  className,
}: React.HTMLAttributes<HTMLElement>) {
  const currentUrl = location.pathname;

  const routes = [
    {
      name: "Dashboard",
      path: "/dashboard",
    },
    {
      name: "Sliders",
      path: "/sliders",
    },
    {
      name: "Categories",
      path: "/categories",
    },
    {
      name: "Products",
      path: "/products",
    },
    {
      name: "Settings",
      path: "/settings",
    },
  ];

  return (
    <nav className={cn("flex items-center space-x-4 lg:space-x-6", className)}>
      {routes.map((route) => (
        <a
          key={route.path}
          href={route.path}
          className={cn(
            "text-sm font-medium text-muted-foreground",
            currentUrl.includes(route.path) ? "font-bold text-secondary" : ""
          )}
        >
          {route.name}
        </a>
      ))}
    </nav>
  );
}
