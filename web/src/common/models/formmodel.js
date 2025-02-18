import {ref} from "vue";

const formData = ref({
    id: 0,
    applicant_type: 'student' | 'employee', // 默认选择学生
    name: '',
    phone: '',
    email: '',
    purposes: '',
    notes: '',
    student: {
        school: '',
        major: '',
        enrollmentYear: '',
        generateYear: '',
    },
    employee: {
        company: '',
        position: '',
        industry: '',
        workExperience: null
    }
});

export default  formData ;