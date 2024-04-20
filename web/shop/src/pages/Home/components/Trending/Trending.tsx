import ProductCard from "@/components/product-card";
import { Product } from "@/interface/Product";
import { ShellIcon } from "lucide-react";

export default function Trending() {
  const trendingProducts: Product[] = [
    {
      id: 1,
      name: "Sdoyri Zufmbk",
      category_id: 1,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 862,
      count_in_stock: 7,
      brand: "erlvxv",
      weight: 753,
      rating: 3,
      slug: "sdoyri-zufmbk",
      image_path: "https://picsum.photos/id/66/720",
      created_at: "2024-04-18 19:25:26.37992 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.37992 +0700 WIB",
      category: {
        id: 1,
        name: "Ceoxwl Aiwujz",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "ceoxwl-aiwujz",
        image_path: "https://picsum.photos/id/12/32/32",
        created_at: "2024-04-18 19:25:26.375833 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.375833 +0700 WIB",
      },
    },
    {
      id: 2,
      name: "Ffwddl Yibeca",
      category_id: 2,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 342,
      count_in_stock: 5,
      brand: "okcung",
      weight: 701,
      rating: 3,
      slug: "ffwddl-yibeca",
      image_path: "https://picsum.photos/id/26/720",
      created_at: "2024-04-18 19:25:26.451505 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.451505 +0700 WIB",
      category: {
        id: 2,
        name: "Qzskrf Gyvxqb",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "qzskrf-gyvxqb",
        image_path: "https://picsum.photos/id/78/32/32",
        created_at: "2024-04-18 19:25:26.448815 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.448815 +0700 WIB",
      },
    },
    {
      id: 3,
      name: "Zwcoyq Ozgcps",
      category_id: 3,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 259,
      count_in_stock: 8,
      brand: "wmtswm",
      weight: 585,
      rating: 3,
      slug: "zwcoyq-ozgcps",
      image_path: "https://picsum.photos/id/36/720",
      created_at: "2024-04-18 19:25:26.521577 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.521577 +0700 WIB",
      category: {
        id: 3,
        name: "Lgihzu Bsqrps",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "lgihzu-bsqrps",
        image_path: "https://picsum.photos/id/27/32/32",
        created_at: "2024-04-18 19:25:26.518412 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.518412 +0700 WIB",
      },
    },
    {
      id: 4,
      name: "Udyliv Mmmlma",
      category_id: 14,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 962,
      count_in_stock: 6,
      brand: "vadgnj",
      weight: 832,
      rating: 3,
      slug: "udyliv-mmmlma",
      image_path: "https://picsum.photos/id/57/720",
      created_at: "2024-04-18 19:25:26.573447 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.573447 +0700 WIB",
      category: {
        id: 14,
        name: "Szqasa Atquph",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "szqasa-atquph",
        image_path: "https://picsum.photos/id/21/32/32",
        created_at: "2024-04-18 19:25:26.571107 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.571107 +0700 WIB",
      },
    },
    {
      id: 5,
      name: "Clmwzk Fcdoab",
      category_id: 15,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 343,
      count_in_stock: 9,
      brand: "hjizpn",
      weight: 573,
      rating: 4,
      slug: "clmwzk-fcdoab",
      image_path: "https://picsum.photos/id/51/720",
      created_at: "2024-04-18 19:25:26.578446 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.578446 +0700 WIB",
      category: {
        id: 15,
        name: "Csapnl Xtipqn",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "csapnl-xtipqn",
        image_path: "https://picsum.photos/id/55/32/32",
        created_at: "2024-04-18 19:25:26.576175 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.576175 +0700 WIB",
      },
    },
    {
      id: 6,
      name: "Agplbj Ssraon",
      category_id: 15,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 620,
      count_in_stock: 7,
      brand: "kceeic",
      weight: 860,
      rating: 3,
      slug: "agplbj-ssraon",
      image_path: "https://picsum.photos/id/97/720",
      created_at: "2024-04-18 19:25:26.581266 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.581266 +0700 WIB",
      category: {
        id: 15,
        name: "Csapnl Xtipqn",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "csapnl-xtipqn",
        image_path: "https://picsum.photos/id/55/32/32",
        created_at: "2024-04-18 19:25:26.576175 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.576175 +0700 WIB",
      },
    },
    {
      id: 7,
      name: "Yzwess Tpzlng",
      category_id: 15,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 331,
      count_in_stock: 3,
      brand: "medcnb",
      weight: 770,
      rating: 3,
      slug: "yzwess-tpzlng",
      image_path: "https://picsum.photos/id/40/720",
      created_at: "2024-04-18 19:25:26.584023 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.584023 +0700 WIB",
      category: {
        id: 15,
        name: "Csapnl Xtipqn",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "csapnl-xtipqn",
        image_path: "https://picsum.photos/id/55/32/32",
        created_at: "2024-04-18 19:25:26.576175 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.576175 +0700 WIB",
      },
    },
    {
      id: 8,
      name: "Trgugw Kngjqn",
      category_id: 15,
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit. Nullam suspendisse auctor diam nascetur magnis cubilia accumsan metus natoque nunc aliquet ipsum luctus ad lacus consequat sodales eu pellentesque fames nisl laoreet porttitor quisque adipiscing est fermentum sed dictumst tincidunt varius torquent potenti tristique. Cubilia purus aptent euismod luctus scelerisque felis volutpat curabitur velit maecenas, habitant sit ac arcu integer erat a lectus dui cursus sapien lorem nisi placerat varius parturient risus morbi aenean porttitor. Ex semper taciti egestas urna ullamcorper sollicitudin mi risus ut laoreet turpis faucibus iaculis sagittis nisl eget suspendisse nec feugiat platea enim vehicula est elementum vivamus etiam. Vulputate sed in iaculis dolor purus tortor adipiscing diam per ultrices bibendum natoque pellentesque nec hendrerit nostra himenaeos suspendisse ad luctus tellus.",
      price: 878,
      count_in_stock: 1,
      brand: "surzhz",
      weight: 877,
      rating: 4,
      slug: "trgugw-kngjqn",
      image_path: "https://picsum.photos/id/34/720",
      created_at: "2024-04-18 19:25:26.58654 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.58654 +0700 WIB",
      category: {
        id: 15,
        name: "Csapnl Xtipqn",
        description:
          "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
        slug: "csapnl-xtipqn",
        image_path: "https://picsum.photos/id/55/32/32",
        created_at: "2024-04-18 19:25:26.576175 +0700 WIB",
        updated_at: "2024-04-18 19:25:26.576175 +0700 WIB",
      },
    },
  ];

  return (
    <div className="space-y-2">
      <div className="w-full border-b-2 pb-1 flex items-center gap-2">
        <ShellIcon className="w-6 h-6" />
        <h1 className="text-2xl font-bold">Trending Products</h1>
      </div>
      <div className="grid grid-cols-4 gap-4">
        {trendingProducts.map((product) => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </div>
  );
}
