<script lang="ts" setup>
  const content = computed(() => ['code', 'games', 'user'])

  const { t } = useI18n({
    useScope: 'local',
  })
</script>

<template>
  <hoshi-no-sora>
    <header-makoto />
    <map-makoto />
    <map-block-with-line v-for="s in content">
      <map-glow-block>
        <template #title> {{ t(s + '.title') }} </template>
        <template #content> {{ t(s + '.content') }} </template>
      </map-glow-block>
    </map-block-with-line>
  </hoshi-no-sora>
</template>

<i18n lang="json" scoped>
{
  "en": {
    "code": {
      "title": "CodeSpace",
      "content": "HeapOverflow"
    },
    "games": {
      "title": "Games",
      "content": "Chess, poker, tetris"
    },
    "user": {
      "title": "User",
      "content": "Admin"
    }
  },
  "ru": {
    "code": {
      "title": "CodeSpace",
      "content": "HeapOverflow"
    },
    "games": {
      "title": "Игры",
      "content": "Шахматы, покер, тетрис"
    },
    "user": {
      "title": "Юзер",
      "content": "Админ"
    }
  }
}
</i18n>
