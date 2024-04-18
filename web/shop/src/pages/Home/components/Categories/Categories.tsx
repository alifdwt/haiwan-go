import { Category } from "@/interface/Category";

export default function Categories() {
  const categories: Category[] = [
    {
      id: 1,
      name: "Ceoxwl Aiwujz",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "ceoxwl-aiwujz",
      image_path: "https://picsum.photos/id/12/32/32",
      created_at: "2024-04-18 19:25:26.375833 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.375833 +0700 WIB",
    },
    {
      id: 2,
      name: "Qzskrf Gyvxqb",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "qzskrf-gyvxqb",
      image_path: "https://picsum.photos/id/78/32/32",
      created_at: "2024-04-18 19:25:26.448815 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.448815 +0700 WIB",
    },
    {
      id: 3,
      name: "Lgihzu Bsqrps",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "lgihzu-bsqrps",
      image_path: "https://picsum.photos/id/27/32/32",
      created_at: "2024-04-18 19:25:26.518412 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.518412 +0700 WIB",
    },
    {
      id: 4,
      name: "Miqjrr Qvmqqh",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "miqjrr-qvmqqh",
      image_path: "https://picsum.photos/id/61/32/32",
      created_at: "2024-04-18 19:25:26.536728 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.536728 +0700 WIB",
    },
    {
      id: 5,
      name: "Vuunmm Scfdzj",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "vuunmm-scfdzj",
      image_path: "https://picsum.photos/id/39/32/32",
      created_at: "2024-04-18 19:25:26.539428 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.539428 +0700 WIB",
    },
    {
      id: 6,
      name: "Qrgebd Uubhpm",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "qrgebd-uubhpm",
      image_path: "https://picsum.photos/id/32/32/32",
      created_at: "2024-04-18 19:25:26.541991 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.541991 +0700 WIB",
    },
    {
      id: 7,
      name: "Plreel Wudzwn",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "plreel-wudzwn",
      image_path: "https://picsum.photos/id/51/32/32",
      created_at: "2024-04-18 19:25:26.544604 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.544604 +0700 WIB",
    },
    {
      id: 8,
      name: "Ybdymr Lkrjfs",
      description:
        "Lorem ipsum dolor sit amet consectetur adipiscing elit proin condimentum venenatis libero nostra lacinia congue viverra duis dui blandit turpis in sapien.",
      slug: "ybdymr-lkrjfs",
      image_path: "https://picsum.photos/id/79/32/32",
      created_at: "2024-04-18 19:25:26.546919 +0700 WIB",
      updated_at: "2024-04-18 19:25:26.546919 +0700 WIB",
    },
  ];

  return (
    <div className="flex items-center justify-between">
      {categories.map((category) => (
        <a
          href={`/categories/${category.slug}`}
          key={category.id}
          className="flex flex-col items-center gap-2 border border-gray-300 rounded-lg p-4 w-36"
        >
          <img
            className="object-cover"
            src={category.image_path}
            alt={category.name}
          />
          <p className="truncate">{category.name}</p>
        </a>
      ))}
    </div>
  );
}
