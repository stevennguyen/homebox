<template>
  <NuxtLink
    class="group card bg-neutral text-neutral-content hover:bg-primary transition-colors duration-300"
    :to="`/item/${item.id}`"
  >
    <div class="card-body py-4 px-6">
      <h2 class="card-title">
        <Icon name="mdi-package-variant" />
        {{ item.name }}
        <Icon v-if="item.archived" class="ml-auto" name="mdi-archive-outline" />
      </h2>
      <p>{{ description }}</p>
      <div class="flex gap-2 flex-wrap justify-end">
        <LabelChip
          v-for="label in item.labels"
          :key="label.id"
          :label="label"
          class="badge-primary group-hover:badge-secondary"
        />
      </div>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
  import { ItemOut, ItemSummary } from "~~/lib/api/types/data-contracts";
  import { truncate } from "~~/lib/strings";

  const props = defineProps({
    item: {
      type: Object as () => ItemOut | ItemSummary,
      required: true,
    },
  });
  const description = computed(() => {
    return truncate(props.item.description, 80);
  });
</script>
