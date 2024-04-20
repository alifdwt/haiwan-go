import { cn } from "@/lib/utils";
import { MoveRightIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";

export default function TransactionBreadcrumb() {
  const currentUrl = location.pathname;
  const navigate = useNavigate();

  const steps = [
    {
      name: "Keranjang",
      url: "/cart",
    },
    {
      name: "Pengiriman",
      url: "/shipment",
    },
    {
      name: "Bayar",
      url: "/payment",
    },
  ];

  return (
    <div className="flex items-center justify-center gap-2">
      {steps.map((step, index) => (
        <div key={index} className="flex items-center gap-1">
          <div
            className="flex items-center gap-1"
            onClick={() => navigate(step.url)}
            role="button"
          >
            <span
              className={cn(
                "flex items-center justify-center w-8 h-8 text-white rounded",
                currentUrl.includes(step.url) ? "bg-primary" : "bg-gray-200"
              )}
            >
              {index + 1}
            </span>
            <p
              className={
                currentUrl.includes(step.url) ? "text-primary" : "text-gray-400"
              }
            >
              {step.name}
            </p>
          </div>
          {index < steps.length - 1 && (
            <MoveRightIcon className="h-5 w-5 text-gray-400" />
          )}
        </div>
      ))}
    </div>
  );
}
