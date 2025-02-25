<template>
  <div v-if="showSelectionColumn" class="bb-grid-cell !px-2">
    <slot name="selection" :database="database" />
  </div>
  <div class="bb-grid-cell">
    <div class="flex items-center space-x-2">
      <SQLEditorButtonV1
        :database="database"
        :disabled="!allowQuery"
        :tooltip="true"
        @failed="$emit('goto-sql-editor-failed')"
      />
      <DatabaseV1Name :database="database" :link="false" tag="span" />
      <BBBadge
        v-if="isPITRDatabaseV1(database)"
        text="PITR"
        :can-remove="false"
        class="text-xs"
      />
      <NTooltip
        v-if="!showMiscColumn && database.syncState !== State.ACTIVE"
        placement="right"
      >
        <template #trigger>
          <heroicons-outline:exclamation-circle class="w-5 h-5 text-error" />
        </template>

        <div class="whitespace-nowrap">
          {{
            $t("database.last-sync-status-long", [
              "NOT_FOUND",
              humanizeDate(database.successfulSyncTime),
            ])
          }}
        </div>
      </NTooltip>
    </div>
  </div>
  <div v-if="showEnvironmentColumn" class="bb-grid-cell">
    <EnvironmentV1Name
      :environment="environment ?? database.effectiveEnvironmentEntity"
      :link="false"
      tag="div"
    />
  </div>
  <div v-if="showSchemaVersionColumn" class="hidden lg:bb-grid-cell">
    {{ database.schemaVersion }}
  </div>
  <div v-if="showProjectColumn" class="bb-grid-cell">
    <ProjectCol
      :project="database.projectEntity"
      :mode="mode"
      :show-tenant-icon="showTenantIcon"
    />
  </div>
  <div v-if="showInstanceColumn" class="bb-grid-cell">
    <InstanceV1Name
      :instance="database.instanceEntity"
      :link="false"
      tag="div"
    />
  </div>
  <div v-if="showMiscColumn" class="bb-grid-cell">
    <div class="w-full flex justify-center">
      <NTooltip placement="left">
        <template #trigger>
          <div
            class="flex items-center justify-center rounded-full select-none w-5 h-5 overflow-hidden text-white font-medium text-base"
            :class="
              database.syncState === State.ACTIVE ? 'bg-success' : 'bg-error'
            "
          >
            <template v-if="database.syncState === State.ACTIVE">
              <heroicons-solid:check class="w-4 h-4" />
            </template>
            <template v-else>
              <span
                class="h-2 w-2 flex items-center justify-center"
                aria-hidden="true"
                >!</span
              >
            </template>
          </div>
        </template>

        <span>
          <template v-if="database.syncState === State.ACTIVE">
            {{
              $t("database.synced-at", {
                time: humanizeDate(database.successfulSyncTime),
              })
            }}
          </template>
          <template v-else>
            {{
              $t("database.not-found-last-successful-sync-was", {
                time: humanizeDate(database.successfulSyncTime),
              })
            }}
          </template>
        </span>
      </NTooltip>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { InstanceV1Name, EnvironmentV1Name } from "@/components/v2";
import { useEnvironmentV1Store } from "@/store";
import { ComposedDatabase } from "@/types";
import { State } from "@/types/proto/v1/common";
import { isPITRDatabaseV1 } from "@/utils";
import ProjectCol from "./ProjectCol.vue";

const props = defineProps<{
  database: ComposedDatabase;
  mode: string;
  showSelectionColumn: boolean;
  showMiscColumn: boolean;
  showSchemaVersionColumn: boolean;
  showProjectColumn: boolean;
  showTenantIcon: boolean;
  showEnvironmentColumn: boolean;
  showInstanceColumn: boolean;
  allowQuery: boolean;
}>();

defineEmits(["goto-sql-editor-failed"]);

const environment = computed(() => {
  return useEnvironmentV1Store().getEnvironmentByName(
    props.database.environment
  );
});
</script>
