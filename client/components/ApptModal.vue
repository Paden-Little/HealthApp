<script setup lang="ts">
import { initFlowbite } from 'flowbite';
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';

const pid = useCookie('pid');
const patient = ref<Patient | null>();

const props = defineProps({
  provider: {
    type: Object as () => Provider,
  },
  services: String,
});

const newPatientBody = reactive({
  firstname: '',
  lastname: '',
  email: '',
  phone: '',
  language: 'english',
  gender: 'male',
  birth: '',
  allergies: [],
  prescriptions: [],
  password: '',
});

const apptBody = reactive({
  date: '',
  start_time: '09:00:00',
  end_time: '1:30',
  provider: props.provider?.id,
  patient: pid.value,
  service: 1,
  description: '',
});

const loadPatientData = async () => {
  try {
    const user = useAuth().getPatientData();
    if (user) {
      user.then(data => {
        patient.value = data;
      });
    }
  } catch (error) {
    console.error('Failed to load patient data:', error);
    alert('Failed to load patient data');
    navigateTo('/my-health/authentication/login');
  }
};

function formatDate(dtf: any) {
  const date = new Date(dtf);
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`;
}

async function createAppointment() {
  // Format date
  apptBody.date = formatDate(apptBody.date);

  // Calculate end_time
  const [hours, minutes, seconds] = apptBody.start_time.split(':');
  let endTime = new Date();
  endTime.setHours(Number(hours), Number(minutes), Number(seconds));

  endTime.setHours(endTime.getHours() + 1);

  apptBody.end_time = `${endTime.getHours().toString().padStart(2, '0')}:${endTime.getMinutes().toString().padStart(2, '0')}:${endTime.getSeconds().toString().padStart(2, '0')}`;

  console.log(apptBody);

  if (pid.value == null) {
    const resp = await $fetch('/api/patient', {
      method: 'POST',
      body: newPatientBody,
    });
    console.log(resp);
  }

  const resp = await $fetch('/api/appointment', {
    method: 'POST',
    body: apptBody,
  });
  console.log(resp);
}

onMounted(() => {
  initFlowbite();
  loadPatientData();
  if (pid.value == null) {
    console.log(false);
  } else {
    console.log(true);
  }
});
</script>

<template>
  <div
    :id="'apptModal' + provider?.email"
    tabindex="-1"
    aria-hidden="true"
    class="fixed left-0 right-0 top-0 z-50 hidden h-[calc(100%-1rem)] max-h-full w-full items-center justify-center overflow-y-auto overflow-x-hidden md:inset-0"
  >
    <div class="relative max-h-full w-full max-w-2xl p-4">
      <!-- Modal content -->
      <div class="relative rounded-lg bg-white shadow">
        <!-- Modal header -->
        <div class="flex justify-between rounded-t border-b p-4 md:p-5">
          <div class="flex">
            <img
              class="mt-2 w-20 rounded-full"
              src="/img/skillissue.jpg"
              alt=""
            />
            <div class="ms-4 flex flex-col">
              <p class="text-lg font-semibold">
                {{ provider?.firstname }} {{ provider?.lastname }},
                {{ provider?.suffix }}
              </p>
              <p class="font-semibold">
                {{ services }}
              </p>
            </div>
          </div>
          <button
            type="button"
            class="ms-auto inline-flex h-8 w-8 items-center justify-center rounded-lg bg-transparent text-sm text-gray-400 hover:bg-gray-200 hover:text-gray-900"
            :data-modal-hide="'apptModal' + provider?.email"
          >
            <Icon name="ic:round-close" />
            <span class="sr-only">Close modal</span>
          </button>
        </div>
        <!-- Modal body -->
        <div class="mx-32 space-y-4 p-4 md:p-5">
          <p class="font-bold">Appointment Date</p>
          <form class="mx-auto">
            <div class="mt-4 grid md:grid-cols-2 md:gap-6">
              <div>
                <label class="mb-2 block text-sm font-medium text-gray-900"
                  >Select date:</label
                >
                <VueDatePicker
                  v-model="apptBody.date"
                  :min-date="new Date()"
                  :enable-time-picker="false"
                />
              </div>
              <div>
                <label
                  for="time"
                  class="mb-2 block text-sm font-medium text-gray-900"
                  >Select time:</label
                >
                <div class="relative">
                  <div
                    class="pointer-events-none absolute inset-y-0 end-0 top-0 flex items-center pe-3.5"
                  >
                    <Icon class="text-gray-400" name="mdi:clock" />
                  </div>
                  <input
                    type="time"
                    id="time"
                    class="block w-full rounded-lg border border-gray-300 p-2.5 text-sm leading-none text-gray-900 focus:border-blue-500 focus:ring-blue-500"
                    min="09:00"
                    max="18:00"
                    required
                    v-model="apptBody.start_time"
                  />
                </div>
              </div>
            </div>
          </form>
        </div>
        <!-- Modal footer -->
        <div
          class="flex flex-col items-center rounded-b border-t border-gray-200 p-4 md:p-5"
        >
          <form class="mx-auto max-w-md">
            <p class="font-bold">User Info</p>
            <div v-if="pid == null">
              <div class="mt-4 grid md:grid-cols-2 md:gap-6">
                <div class="group relative z-0 mb-5 w-full">
                  <input
                    type="text"
                    name="floating_first_name"
                    id="floating_first_name"
                    class="peer block w-full appearance-none border-0 border-b-2 border-gray-300 bg-transparent px-0 py-2.5 text-sm text-gray-900 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500"
                    placeholder=" "
                    required
                    v-model="newPatientBody.firstname"
                  />
                  <label
                    for="floating_first_name"
                    class="absolute top-3 -z-10 origin-[0] -translate-y-6 scale-75 transform text-sm text-gray-500 duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:start-0 peer-focus:-translate-y-6 peer-focus:scale-75 peer-focus:font-medium peer-focus:text-blue-600 dark:text-gray-400 peer-focus:dark:text-blue-500 rtl:peer-focus:translate-x-1/4"
                    >First name</label
                  >
                </div>
                <div class="group relative z-0 mb-5 w-full">
                  <input
                    type="text"
                    name="floating_last_name"
                    id="floating_last_name"
                    class="peer block w-full appearance-none border-0 border-b-2 border-gray-300 bg-transparent px-0 py-2.5 text-sm text-gray-900 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500"
                    placeholder=" "
                    required
                    v-model="newPatientBody.lastname"
                  />
                  <label
                    for="floating_last_name"
                    class="absolute top-3 -z-10 origin-[0] -translate-y-6 scale-75 transform text-sm text-gray-500 duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:start-0 peer-focus:-translate-y-6 peer-focus:scale-75 peer-focus:font-medium peer-focus:text-blue-600 dark:text-gray-400 peer-focus:dark:text-blue-500 rtl:peer-focus:translate-x-1/4"
                    >Last name</label
                  >
                </div>
              </div>
              <div class="group relative z-0 mb-5 w-full">
                <input
                  type="email"
                  name="floating_email"
                  id="floating_email"
                  class="peer block w-full appearance-none border-0 border-b-2 border-gray-300 bg-transparent px-0 py-2.5 text-sm text-gray-900 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500"
                  placeholder=" "
                  required
                  v-model="newPatientBody.email"
                />
                <label
                  for="floating_email"
                  class="absolute top-3 -z-10 origin-[0] -translate-y-6 scale-75 transform text-sm text-gray-500 duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:start-0 peer-focus:-translate-y-6 peer-focus:scale-75 peer-focus:font-medium peer-focus:text-blue-600 dark:text-gray-400 peer-focus:dark:text-blue-500 rtl:peer-focus:left-auto rtl:peer-focus:translate-x-1/4"
                  >Email address</label
                >
              </div>
              <div class="group relative z-0 mb-5 w-full">
                <input
                  type="tel"
                  pattern="[0-9]{10}"
                  name="floating_phone"
                  id="floating_phone"
                  class="peer block w-full appearance-none border-0 border-b-2 border-gray-300 bg-transparent px-0 py-2.5 text-sm text-gray-900 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500"
                  placeholder=" "
                  required
                  v-model="newPatientBody.phone"
                />
                <label
                  for="floating_phone"
                  class="absolute top-3 -z-10 origin-[0] -translate-y-6 scale-75 transform text-sm text-gray-500 duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:start-0 peer-focus:-translate-y-6 peer-focus:scale-75 peer-focus:font-medium peer-focus:text-blue-600 dark:text-gray-400 peer-focus:dark:text-blue-500 rtl:peer-focus:translate-x-1/4"
                  >Phone number (123-456-7890)</label
                >
              </div>

              <div class="group relative z-0 mb-5 w-full">
                <input
                  type="text"
                  name="floating_password"
                  id="floating_password"
                  class="peer block w-full appearance-none border-0 border-b-2 border-gray-300 bg-transparent px-0 py-2.5 text-sm text-gray-900 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500"
                  placeholder=" "
                  required
                  v-model="newPatientBody.password"
                />
                <label
                  for="floating_password"
                  class="absolute top-3 -z-10 origin-[0] -translate-y-6 scale-75 transform text-sm text-gray-500 duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:start-0 peer-focus:-translate-y-6 peer-focus:scale-75 peer-focus:font-medium peer-focus:text-blue-600 dark:text-gray-400 peer-focus:dark:text-blue-500 rtl:peer-focus:left-auto rtl:peer-focus:translate-x-1/4"
                  >Password</label
                >
              </div>
              <div>
                <label class="mb-2 block text-sm font-medium text-gray-500"
                  >Date of Birth</label
                >
                <VueDatePicker
                  v-model="apptBody.date"
                  :enable-time-picker="false"
                />
              </div>
            </div>
            <div class="group relative z-0 mb-5 w-full">
              <label
                for="large-input"
                class="mb-2 mt-4 block text-sm font-medium text-gray-500 dark:text-white"
                >Reason</label
              >
              <input
                type="text"
                id="large-input"
                v-model="apptBody.description"
                class="block w-full rounded-lg border border-gray-300 p-4 text-base text-gray-900 focus:border-blue-500 focus:ring-blue-500"
              />
            </div>
          </form>
          <div class="flex justify-start">
            <button
              :data-modal-hide="'apptModal' + provider?.email"
              type="button"
              class="cta-btn"
              @click="createAppointment()"
            >
              Book Appointment
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
