<template>
  <div class="space-y-3 w-full overflow-x-auto">
    <div
      class="w-full border-b pb-2 mb-2 flex flex-row justify-between items-center"
    >
      <div class="flex flex-row justify-start items-center space-x-2"></div>
      <div>
        <NButton type="primary" @click="state.showCreatePanel = true">
          <heroicons-solid:plus class="w-4 h-auto mr-0.5" />
          <span>{{ $t("database.new-branch") }}</span>
        </NButton>
      </div>
    </div>

    <BranchTable
      v-if="ready"
      :branches="sortedBranches"
      :hide-project-column="hideProjectColumn"
      @click="handleBranchClick"
    />
    <div v-else class="w-full h-[20rem] flex items-center justify-center">
      <BBSpin />
    </div>
  </div>

  <CreateSchemaDesignPanel
    v-if="state.showCreatePanel"
    :project-id="projectId"
    @dismiss="state.showCreatePanel = false"
    @created="
      (schemaDesign: SchemaDesign) => {
        state.showCreatePanel = false;
        handleBranchClick(schemaDesign);
      }
    "
  />

  <EditSchemaDesignPanel
    v-if="state.selectedSchemaDesignName"
    :schema-design-name="state.selectedSchemaDesignName"
    @dismiss="state.selectedSchemaDesignName = undefined"
  />
</template>

<script lang="ts" setup>
import { orderBy } from "lodash-es";
import { NButton } from "naive-ui";
import { computed, reactive } from "vue";
import BranchTable from "@/components/SchemaDesigner/BranchTable.vue";
import CreateSchemaDesignPanel from "@/components/SchemaDesigner/CreateSchemaDesignPanel.vue";
import EditSchemaDesignPanel from "@/components/SchemaDesigner/EditSchemaDesignPanel.vue";
import { useProjectV1Store } from "@/store";
import { useSchemaDesignList } from "@/store/modules/schemaDesign";
import { SchemaDesign } from "@/types/proto/v1/schema_design_service";

const props = defineProps<{
  projectId?: string;
  hideProjectColumn?: boolean;
}>();

interface LocalState {
  showCreatePanel: boolean;
  selectedSchemaDesignName?: string;
}

const projectV1Store = useProjectV1Store();
const { schemaDesignList, ready } = useSchemaDesignList();
const state = reactive<LocalState>({
  showCreatePanel: false,
});

const project = computed(() =>
  projectV1Store.getProjectByUID(props.projectId || "")
);

const sortedBranches = computed(() => {
  return orderBy(
    props.projectId
      ? schemaDesignList.value.filter((schemaDesign) =>
          schemaDesign.name.startsWith(project.value.name)
        )
      : schemaDesignList.value,
    "updateTime",
    "desc"
  );
});

const handleBranchClick = async (schemaDesign: SchemaDesign) => {
  state.selectedSchemaDesignName = schemaDesign.name;
};
</script>
