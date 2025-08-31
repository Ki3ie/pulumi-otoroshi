import {Organization, Tenant, ServiceGroup} from "@pulumi/otoroshi/organize";
import { Route } from "@pulumi/otoroshi/proxy"

const organizationName = "organization-1";
new Organization(organizationName, {
    name: organizationName,
});

const tenantName = "tenant-1";
new Tenant(tenantName, {
    name: tenantName
});

const serviceGrouoName = "service-group-5";
new ServiceGroup(serviceGrouoName, {
    name: serviceGrouoName,
    location: {
        tenant: "test",
        teams: ["test"]
    }
});

const routeName = "route-1";
new Route(routeName, {
    name: routeName
})
