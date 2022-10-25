import { TableColumn } from "react-data-table-component";
import ResourceView from "../components/ResourceView";
import { extractResourceGroup, extractResourceName, Resource } from "../resources/resources";

interface ContainerProperties {
  application: string
  environment: string
}

const columns: TableColumn<Resource<ContainerProperties>>[] = [
  {
    name: 'Name',
    selector: (row: Resource<ContainerProperties>) => row.name,
    sortable: true,
  },
  {
    name: 'Type',
    selector: (row: Resource<ContainerProperties>) => row.type,
    sortable: true,
  },
  {
    name: 'Resource Group',
    selector: (row: Resource<ContainerProperties>) => extractResourceGroup(row.id) ?? '',
    sortable: true,
  },
  {
    name: 'Application',
    selector: (row: Resource<ContainerProperties>) => extractResourceName(row.properties.application) ?? '',
    sortable: true,
  },
  {
    name: 'Environment',
    selector: (row: Resource<ContainerProperties>) => extractResourceName(row.properties.environment) ?? '',
    sortable: true,
  },
];


export default function ResourcePage() {
  return (
    <>
      <ResourceView
        columns={columns}
        heading="All Resources"
        resourceType={undefined}
        selectionMessage="Select a resource to display details..." />
    </>
  );
}