<script lang="ts" setup>
  import { StarsBackground } from 'makoto-ui-vue'
  const content = computed(() => ['code', 'games', 'user', 'dict', 'chat'])

  const { t } = useI18n({
    useScope: 'local',
  })
</script>

<template>
  <stars-background>
    <header-makoto />
    <map-makoto />
    <map-block-with-line v-for="s in content">
      <map-glow-block>
        <template #title> {{ t(s + '.title') }} </template>
        <template #content> {{ t(s + '.content') }} </template>
      </map-glow-block>
    </map-block-with-line>
  </stars-background>
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
    },
    "dict": {
      "title": "Dictionary",
      "content": "Key-word pairs"
    },
    "chat": {
      "title": "Chat",
      "content": "Makoto Community"
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
    },
    "dict": {
      "title": "Словарик",
      "content": "Пары ключ-значение"
    },
    "chat": {
      "title": "Чат",
      "content": "Комьюнити Макото"
    }
  }
}
</i18n>

<style>
  .stars__content {
    overflow: hidden !important;
  }
</style>
