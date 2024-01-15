---
Title: "Backup Schedule Resource"
Description: |-
   Creating a backup schedule.
---

# Backup Schedule Resource

This resource enables users to create and configure scheduled backups in a cluster or cluster group level.
Backups can be applied in 3 levels:

* Entire Cluster
* Selected Namespaces
* Resources Selection By Label Selector

For more information regarding scheduled backups, see [Scheduled Backups][backup-schedule].

[backup-schedule]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-89926F80-050A-4F1C-9D04-D56D5F453995.html?hWord=N4IghgNiBcIEZgMYGsCuAHABAZ0QCwFMATVCAkAXyA

**NOTE :** To resolve cluster and cluster group backup schedule conflicts use the below command
``terraform refresh``

For instance, in case cluster group/cluster data protection is disabled then use above command
and remove cluster group/cluster backup schedule resource from terraform file.

# Entire Cluster Weekly Backup Schedule

## Example Usage

```terraform
resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name = "full-weekly"
  scope {
    cluster {
      management_cluster_name = "MGMT_CLS_NAME"
      provisioner_name        = "PROVISIONER_NAME"
      cluster_name            = "CLS_NAME"
    }
  }

  backup_scope = "FULL_CLUSTER"

  spec {
    schedule {
      rate = "0 12 * * 1"
    }

    template {
      backup_ttl = "2592000s"
      excluded_namespaces = [
        "app-01",
        "app-02",
        "app-03",
        "app-04"
      ]
      excluded_resources = [
        "secrets",
        "configmaps"
      ]

      storage_location = "TARGET_LOCATION_NAME"
    }
  }
}
```


# Selected Namespaces Hourly Backup Schedule

## Example Usage

```terraform
resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name = "namespaces-hourly"
  scope {
    cluster {
      management_cluster_name = "MGMT_CLS_NAME"
      provisioner_name        = "PROVISIONER_NAME"
      cluster_name            = "CLS_NAME"
    }
  }

  backup_scope = "SET_NAMESPACES"

  spec {
    schedule {
      rate = "30 * * * *"
    }

    template {
      included_namespaces = [
        "app-01",
        "app-02",
        "app-03",
        "app-04"
      ]

      excluded_resources = [
        "secrets",
        "configmaps"
      ]

      backup_ttl                = "86400s"
      include_cluster_resources = true
      storage_location          = "TARGET_LOCATION_NAME"

      hooks {
        resource {
          name = "sample-config"

          pre_hook {
            exec {
              command   = ["echo 'hello'"]
              container = "workload"
              on_error  = "CONTINUE"
              timeout   = "10s"
            }
          }

          pre_hook {
            exec {
              command   = ["echo 'hello'"]
              container = "db"
              on_error  = "CONTINUE"
              timeout   = "30s"
            }
          }

          post_hook {
            exec {
              command   = ["echo 'goodbye'"]
              container = "db"
              on_error  = "CONTINUE"
              timeout   = "60s"
            }
          }

          post_hook {
            exec {
              command   = ["echo 'goodbye'"]
              container = "workload"
              on_error  = "FAIL"
              timeout   = "20s"
            }
          }
        }
      }
    }
  }
}
```

# Resources Selection By Label Selector Backup Schedule

## Example Usage

```terraform
resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name = "label-based-no-storage"
  scope {
    cluster {
      management_cluster_name = "MGMT_CLS_NAME"
      provisioner_name        = "PROVISIONER_NAME"
      cluster_name            = "CLS_NAME"
    }
  }

  backup_scope = "LABEL_SELECTOR"


  spec {
    schedule {
      rate = "0 12 * * *"
    }

    template {
      default_volumes_to_fs_backup = false
      include_cluster_resources    = true
      backup_ttl                   = "604800s"
      storage_location             = "TARGET_LOCATION_NAME"

      label_selector {
        match_expression {
          key      = "apps.tanzu.vmware.com/demo"
          operator = "Exists"
        }

        match_expression {
          key      = "apps.tanzu.vmware.com/exclude-from-backup"
          operator = "DoesNotExist"
        }
      }
    }
  }
}
```


# Entire Cluster Group Weekly Backup Schedule

## Example Usage

```terraform
resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name = "full-weekly"
  scope {
    cluster_group {
      cluster_group_name = "CG_NAME"
    }
  }
  selector {
    names = [
      "cluster1",
      "cluster2"
    ]
  }

  backup_scope = "FULL_CLUSTER"
  spec {
    schedule {
      rate = "0 12 * * 1"
    }

    template {
      backup_ttl = "2592000s"
      excluded_namespaces = [
        "app-01",
        "app-02",
        "app-03",
        "app-04"
      ]
      excluded_resources = [
        "secrets",
        "configmaps"
      ]

      storage_location = "TARGET_LOCATION_NAME"
    }
  }
}
```


# Selected Namespaces Hourly Cluster Group Backup Schedule

## Example Usage

```terraform
resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name = "namespaces-hourly"
  scope {
    cluster_group {
      cluster_group_name = "CG_NAME"
    }
  }
  selector {
    names = [
      "cluster1",
      "cluster2"
    ]
  }

  backup_scope = "SET_NAMESPACES"

  spec {
    schedule {
      rate = "30 * * * *"
    }

    template {
      included_namespaces = [
        "app-01",
        "app-02",
        "app-03",
        "app-04"
      ]

      excluded_resources = [
        "secrets",
        "configmaps"
      ]

      backup_ttl                = "86400s"
      include_cluster_resources = true
      storage_location          = "TARGET_LOCATION_NAME"

      hooks {
        resource {
          name = "sample-config"

          pre_hook {
            exec {
              command   = ["echo 'hello'"]
              container = "workload"
              on_error  = "CONTINUE"
              timeout   = "10s"
            }
          }

          pre_hook {
            exec {
              command   = ["echo 'hello'"]
              container = "db"
              on_error  = "CONTINUE"
              timeout   = "30s"
            }
          }

          post_hook {
            exec {
              command   = ["echo 'goodbye'"]
              container = "db"
              on_error  = "CONTINUE"
              timeout   = "60s"
            }
          }

          post_hook {
            exec {
              command   = ["echo 'goodbye'"]
              container = "workload"
              on_error  = "FAIL"
              timeout   = "20s"
            }
          }
        }
      }
    }
  }
}
```

# Resources Selection By Label Selector Cluster Group Backup Schedule

## Example Usage

```terraform
resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name = "label-based-no-storage"
  scope {
    cluster_group {
      cluster_group_name = "CG_NAME"
    }
  }
  selector {
    names = [
      "cluster1",
      "cluster2"
    ]
  }
  backup_scope = "LABEL_SELECTOR"


  spec {
    schedule {
      rate = "0 12 * * *"
    }

    template {
      default_volumes_to_fs_backup = false
      include_cluster_resources    = true
      backup_ttl                   = "604800s"
      storage_location             = "TARGET_LOCATION_NAME"

      label_selector {
        match_expression {
          key      = "apps.tanzu.vmware.com/demo"
          operator = "Exists"
        }

        match_expression {
          key      = "apps.tanzu.vmware.com/exclude-from-backup"
          operator = "DoesNotExist"
        }
      }
    }
  }
}
```

## Import Backup Schedule
The resource ID for importing an existing backup schedule should be comprised of a full cluster name and a backup schedule name separated by '/'.

```bash
terraform import tanzu-mission-control_backup_schedule.demo_backup MANAGEMENT_CLUSTER_NAME/PROVISIONER_NAME/CLUSTER_NAME/BACKUP_SCHEDULE_NAME
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `backup_scope` (String) Scope for backup schedule.
Valid values are (FULL_CLUSTER, SET_NAMESPACES, LABEL_SELECTOR)
- `name` (String) The name of the backup schedule
- `scope` (Block List, Min: 1, Max: 1) Scope block for Back up schedule (cluster/cluster group) (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Backup schedule spec block (see [below for nested schema](#nestedblock--spec))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `selector` (Block List) Selector of the cluster group backup schedule (see [below for nested schema](#nestedblock--selector))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `cluster` (Block List, Max: 1) Cluster scope block (see [below for nested schema](#nestedblock--scope--cluster))
- `cluster_group` (Block List, Max: 1) Cluster group scope block (see [below for nested schema](#nestedblock--scope--cluster_group))

<a id="nestedblock--scope--cluster"></a>
### Nested Schema for `scope.cluster`

Required:

- `cluster_name` (String) Cluster name
- `management_cluster_name` (String) Management cluster name
- `provisioner_name` (String) Cluster provisioner name


<a id="nestedblock--scope--cluster_group"></a>
### Nested Schema for `scope.cluster_group`

Required:

- `cluster_group_name` (String) Cluster group name



<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `schedule` (Block List, Min: 1, Max: 1) Schedule block (see [below for nested schema](#nestedblock--spec--schedule))

Optional:

- `paused` (Boolean) Paused specifies whether the schedule is paused or not. (Default: False)
- `template` (Block List, Max: 1) Backup schedule template block, backup definition to be run on the provided schedule (see [below for nested schema](#nestedblock--spec--template))

<a id="nestedblock--spec--schedule"></a>
### Nested Schema for `spec.schedule`

Required:

- `rate` (String) Cron expression of backup schedule rate/interval


<a id="nestedblock--spec--template"></a>
### Nested Schema for `spec.template`

Optional:

- `backup_ttl` (String) The backup retention period.
- `csi_snapshot_timeout` (String) Specifies the time used to wait for CSI VolumeSnapshot status turns to ReadyToUse during creation, before returning error as timeout.
The default value is 10 minute.
Format is the time number and time sign, example: "50s" (50 seconds)
- `default_volumes_to_fs_backup` (Boolean) Specifies whether all pod volumes should be backed up via file system backup by default.
(Default: True)
- `default_volumes_to_restic` (Boolean) Specifies whether restic should be used to take a backup of all pod volumes by default.
(Default: False)
- `excluded_cluster_scoped_resources` (List of String) List of cluster-scoped resource type names to exclude from the backup.
If set to "*", all cluster-scoped resource types are excluded.
- `excluded_namespace_scoped_resources` (List of String) List of of namespace-scoped resource type names to exclude from the backup.
If set to "*", all namespace-scoped resource types are excluded.
- `excluded_namespaces` (List of String) The namespaces to be excluded in the backup.
Can't be used if scope is SET_NAMESPACES.
- `excluded_resources` (List of String) The name list for the resources to be excluded in backup.
- `hooks` (Block List, Max: 1) Hooks block represent custom actions that should be executed at different phases of the backup. (see [below for nested schema](#nestedblock--spec--template--hooks))
- `include_cluster_resources` (Boolean) A flag which specifies whether cluster-scoped resources should be included for consideration in the backup.
If set to true, all cluster-scoped resources will be backed up. If set to false, all cluster-scoped resources will be excluded from the backup.
If unset, all cluster-scoped resources are included if and only if all namespaces are included and there are no excluded namespaces.
Otherwise, only cluster-scoped resources associated with namespace-scoped resources included in the backup spec are backed up.
For example, if a PersistentVolumeClaim is included in the backup, its associated PersistentVolume (which is cluster-scoped) would also be backed up.
(Default: False)
- `included_cluster_scoped_resources` (List of String) List of cluster-scoped resource type names to include in the backup.
If set to "*", all cluster-scoped resource types are included.
  The default value is empty, which means only related cluster-scoped resources are included.
- `included_namespace_scoped_resources` (List of String) List of of namespace-scoped resource type names to include in the backup.
The default value is "*".
- `included_namespaces` (List of String) The namespace to be included for backup from.
If empty, all namespaces are included.
Can't be used if scope is FULL_CLUSTER.
Required if scope is SET_NAMESPACES.
- `included_resources` (List of String) The name list for the resources to be included into backup. If empty, all resources are included.
- `label_selector` (Block List, Max: 1) The label selector to selectively adding individual objects to the backup schedule.
If not specified, all objects are included.
Can't be used if scope is FULL_CLUSTER or SET_NAMESPACES.
Required if scope is LABEL_SELECTOR and Or Label Selectors are not defined (see [below for nested schema](#nestedblock--spec--template--label_selector))
- `or_label_selector` (Block List) (Repeatable Block) A list of label selectors to filter with when adding individual objects to the backup.
If multiple provided they will be joined by the OR operator.
LabelSelector as well as OrLabelSelectors cannot co-exist in backup request, only one of them can be used.
Can't be used if scope is FULL_CLUSTER or SET_NAMESPACES.
Required if scope is LABEL_SELECTOR and Label Selector is not defined (see [below for nested schema](#nestedblock--spec--template--or_label_selector))
- `ordered_resources` (Map of String) Specifies the backup order of resources of specific Kind. The map key is the Kind name and value is a list of resource names separated by commas.
Each resource name has format "namespace/resourcename".
For cluster resources, simply use "resourcename".
- `snapshot_move_data` (Boolean) Specifies whether snapshot data should be moved to the target location.(Default:False)
- `snapshot_volumes` (Boolean) A flag which specifies whether to take cloud snapshots of any PV's referenced in the set of objects included in the Backup.
If set to true, snapshots will be taken, otherwise, snapshots will be skipped.
If left unset, snapshots will be attempted if volume snapshots are configured for the cluster.
- `storage_location` (String) The name of a BackupStorageLocation where the backup should be stored.
- `volume_snapshot_locations` (List of String) A list containing names of VolumeSnapshotLocations associated with this backup.

Read-Only:

- `sys_excluded_namespaces` (List of String) System excluded namespaces for state.

<a id="nestedblock--spec--template--hooks"></a>
### Nested Schema for `spec.template.hooks`

Optional:

- `resource` (Block List) (Repeatable Block) Resources are hooks that should be executed when backing up individual instances of a resource. (see [below for nested schema](#nestedblock--spec--template--hooks--resource))

<a id="nestedblock--spec--template--hooks--resource"></a>
### Nested Schema for `spec.template.hooks.resource`

Required:

- `name` (String) The name of the hook resource.

Optional:

- `excluded_namespaces` (List of String) Specifies the namespaces to which this hook spec does not apply.
- `included_namespaces` (List of String) Specifies the namespaces to which this hook spec applies.
If empty, it applies to all namespaces.
- `label_selector` (Block List, Max: 1) The label selector to selectively adding individual objects to the hook resource.
If not specified, all objects are included. (see [below for nested schema](#nestedblock--spec--template--hooks--resource--label_selector))
- `post_hook` (Block List) (Repeatable Block) A list of backup hooks to execute after storing the item in the backup.
These are executed after all "additional items" from item actions are processed. (see [below for nested schema](#nestedblock--spec--template--hooks--resource--post_hook))
- `pre_hook` (Block List) (Repeatable Block) A list of backup hooks to execute after storing the item in the backup.
These are executed after all "additional items" from item actions are processed. (see [below for nested schema](#nestedblock--spec--template--hooks--resource--pre_hook))

<a id="nestedblock--spec--template--hooks--resource--label_selector"></a>
### Nested Schema for `spec.template.hooks.resource.label_selector`

Optional:

- `match_expression` (Block List) (Repeatable Block) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--hooks--resource--label_selector--match_expression))
- `match_labels` (Map of String) A map of {key,value} pairs. A single {key,value} in the map is equivalent to an element of match_expressions, whose key field is "key", the operator is "In" and the values array contains only "value".
The requirements are ANDed.

<a id="nestedblock--spec--template--hooks--resource--label_selector--match_expression"></a>
### Nested Schema for `spec.template.hooks.resource.label_selector.match_expression`

Required:

- `key` (String) Key is the label key that the selector applies to.
- `operator` (String) Operator represents a key's relationship to a set of values.
Valid operators are "In", "NotIn", "Exists" and "DoesNotExist".

Optional:

- `values` (List of String) Values is an array of string values.
If the operator is "In" or "NotIn", the values array must be non-empty.
If the operator is "Exists" or "DoesNotExist", the values array must be empty.
This array is replaced during a strategic merge patch.



<a id="nestedblock--spec--template--hooks--resource--post_hook"></a>
### Nested Schema for `spec.template.hooks.resource.post_hook`

Required:

- `exec` (Block List, Min: 1, Max: 1) Exec block defines an exec hook. (see [below for nested schema](#nestedblock--spec--template--hooks--resource--post_hook--exec))

<a id="nestedblock--spec--template--hooks--resource--post_hook--exec"></a>
### Nested Schema for `spec.template.hooks.resource.post_hook.exec`

Required:

- `command` (List of String) The command and arguments to execute.
- `container` (String) The container in the pod where the command should be executed.
If not specified, the pod's first container is used.

Optional:

- `on_error` (String) Specifies how Velero should behave if it encounters an error executing this hook.
Valid values are (FAIL, CONTINUE)
- `timeout` (String) Defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.



<a id="nestedblock--spec--template--hooks--resource--pre_hook"></a>
### Nested Schema for `spec.template.hooks.resource.pre_hook`

Required:

- `exec` (Block List, Min: 1, Max: 1) Exec block defines an exec hook. (see [below for nested schema](#nestedblock--spec--template--hooks--resource--pre_hook--exec))

<a id="nestedblock--spec--template--hooks--resource--pre_hook--exec"></a>
### Nested Schema for `spec.template.hooks.resource.pre_hook.exec`

Required:

- `command` (List of String) The command and arguments to execute.
- `container` (String) The container in the pod where the command should be executed.
If not specified, the pod's first container is used.

Optional:

- `on_error` (String) Specifies how Velero should behave if it encounters an error executing this hook.
Valid values are (FAIL, CONTINUE)
- `timeout` (String) Defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.





<a id="nestedblock--spec--template--label_selector"></a>
### Nested Schema for `spec.template.label_selector`

Optional:

- `match_expression` (Block List) (Repeatable Block) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--label_selector--match_expression))
- `match_labels` (Map of String) A map of {key,value} pairs. A single {key,value} in the map is equivalent to an element of match_expressions, whose key field is "key", the operator is "In" and the values array contains only "value".
The requirements are ANDed.

<a id="nestedblock--spec--template--label_selector--match_expression"></a>
### Nested Schema for `spec.template.label_selector.match_expression`

Required:

- `key` (String) Key is the label key that the selector applies to.
- `operator` (String) Operator represents a key's relationship to a set of values.
Valid operators are "In", "NotIn", "Exists" and "DoesNotExist".

Optional:

- `values` (List of String) Values is an array of string values.
If the operator is "In" or "NotIn", the values array must be non-empty.
If the operator is "Exists" or "DoesNotExist", the values array must be empty.
This array is replaced during a strategic merge patch.



<a id="nestedblock--spec--template--or_label_selector"></a>
### Nested Schema for `spec.template.or_label_selector`

Optional:

- `match_expression` (Block List) (Repeatable Block) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--or_label_selector--match_expression))
- `match_labels` (Map of String) A map of {key,value} pairs. A single {key,value} in the map is equivalent to an element of match_expressions, whose key field is "key", the operator is "In" and the values array contains only "value".
The requirements are ANDed.

<a id="nestedblock--spec--template--or_label_selector--match_expression"></a>
### Nested Schema for `spec.template.or_label_selector.match_expression`

Required:

- `key` (String) Key is the label key that the selector applies to.
- `operator` (String) Operator represents a key's relationship to a set of values.
Valid operators are "In", "NotIn", "Exists" and "DoesNotExist".

Optional:

- `values` (List of String) Values is an array of string values.
If the operator is "In" or "NotIn", the values array must be non-empty.
If the operator is "Exists" or "DoesNotExist", the values array must be empty.
This array is replaced during a strategic merge patch.





<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--selector"></a>
### Nested Schema for `selector`

Optional:

- `excluded_names` (List of String) Specifies the name of excluded clusters.
- `label_selector` (Block List, Max: 1) The label selector to selectively adding individual clusters to the cluster group backup schedule.
If not specified, all clusters are included. (see [below for nested schema](#nestedblock--selector--label_selector))
- `names` (List of String) Specifies name of cluster to be selected.

<a id="nestedblock--selector--label_selector"></a>
### Nested Schema for `selector.label_selector`

Optional:

- `match_expression` (Block List) (Repeatable Block) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--selector--label_selector--match_expression))
- `match_labels` (Map of String) A map of {key,value} pairs. A single {key,value} in the map is equivalent to an element of match_expressions, whose key field is "key", the operator is "In" and the values array contains only "value".
The requirements are ANDed.

<a id="nestedblock--selector--label_selector--match_expression"></a>
### Nested Schema for `selector.label_selector.match_expression`

Required:

- `key` (String) Key is the label key that the selector applies to.
- `operator` (String) Operator represents a key's relationship to a set of values.
Valid operators are "In", "NotIn", "Exists" and "DoesNotExist".

Optional:

- `values` (List of String) Values is an array of string values.
If the operator is "In" or "NotIn", the values array must be non-empty.
If the operator is "Exists" or "DoesNotExist", the values array must be empty.
This array is replaced during a strategic merge patch.
