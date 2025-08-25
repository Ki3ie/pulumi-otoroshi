import {Organization, Tenant} from "@pulumi/otoroshi/organize";

const organizationName = "organization-1";
new Organization(organizationName, {
    name: organizationName,
});

const tenantName = "tenant-1";
new Tenant(tenantName, {
    name: tenantName
});
