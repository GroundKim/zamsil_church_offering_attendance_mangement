<template>
  <v-footer padless class="mt-5">
    <v-card
      flat
      tile
      class="flex text-center"
    >
      <v-card-text>
        <v-container>
          <v-row class="pt-3">
            <v-col class="pt-0">
              <v-chip outlined class="mx-2">
                <v-icon class="mx-1">mdi-card-account-phone</v-icon>
                +82-10-8986-2379
              </v-chip>
            </v-col>
            <v-col class="pt-0">
              <v-chip outlined class="mx-2">
                <v-icon class="mx-1">mdi-at</v-icon>
                jjoo0815@gmail.com
              </v-chip>
            </v-col>
            <v-col class="pt-0">
              <v-chip outlined class="mx-2">
                <strong class="mx-1">KAKAO</strong>
                jjoo0815
              </v-chip>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      
      <v-card-text class="pt-0">
        <strong>&copy;{{ new Date().getFullYear() }} GroundKim</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<style scoped>
.v-footer {
  border-top: 1px solid #f5f5f5;
}
</style>
