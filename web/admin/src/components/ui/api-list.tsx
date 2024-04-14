import { myApi } from "@/helpers/api";
import { ApiAlert } from "./api-alert";

interface ApiListProps {
  entityName: string;
  entityIdName: string;
  superAdmin?: boolean;
}

export const ApiList: React.FC<ApiListProps> = ({
  entityName,
  entityIdName,
  superAdmin,
}) => {
  const baseUrl = myApi.defaults.baseURL;

  return (
    <>
      <ApiAlert
        title="GET"
        variant="public"
        description={`${baseUrl}/${entityName}`}
      />
      <ApiAlert
        title="GET"
        variant="public"
        description={`${baseUrl}/${entityName}/${entityIdName}`}
      />
      {superAdmin && (
        <>
          <ApiAlert
            title="POST"
            variant="admin"
            description={`${baseUrl}/${entityName}`}
          />
          <ApiAlert
            title="PUT"
            variant="admin"
            description={`${baseUrl}/${entityName}/${entityIdName}`}
          />
          <ApiAlert
            title="DELETE"
            variant="admin"
            description={`${baseUrl}/${entityName}/${entityIdName}`}
          />
        </>
      )}
    </>
  );
};
