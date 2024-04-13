import { cn } from "@/lib/utils";
import React from "react";

export function MainNav({ className }: React.HTMLAttributes<HTMLElement>) {
  const currentUrl = location.pathname;
  const routes = [
    {
      href: "/",
      label: "Home",
    },
    {
      href: "/register",
      label: "Register",
    },
    {
      href: "/login",
      label: "Login",
    },
  ];

  return (
    <nav
      className={cn("mx-6 flex items-center space-x-4 lg:space-x-6", className)}
    >
      {routes.map((item) => (
        <a
          key={item.href}
          href={item.href}
          className={cn(
            "text-sm font-medium text-muted-foreground hover:text-white",
            currentUrl === item.href && "text-white"
          )}
        >
          {item.label}
        </a>
      ))}
    </nav>
  );
}
