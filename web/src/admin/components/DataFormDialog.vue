<template>
  <el-dialog
      :model-value="visible"
      :title="isEdit ? '编辑报名者信息' : '新增报名者'"

      @update:model-value="$emit('update:visible', $event)"
      width="600px"
  >
    <el-form
        :model="formData"
        :rule = "formRules"
        ref="formRef"
        label-width="120px">
      <el-form-item label="姓名" prop="name">
        <el-input v-model="formData.name" />
      </el-form-item>

      <el-form-item label="类型" prop="applicant_type">
        <el-radio-group v-model="formData.applicant_type">
          <el-radio label="student">学生</el-radio>
          <el-radio label="employee">在职人员</el-radio>
        </el-radio-group>
      </el-form-item>

      <!-- 学生信息 -->
      <template v-if="formData.applicant_type === 'student'">
        <el-form-item label="学校" prop="student.school">
          <el-input v-model="formData.student.school" />
        </el-form-item>
        <el-form-item label="专业" prop="student.major">
          <el-input v-model="formData.student.major" />
        </el-form-item>
        <el-form-item label="入学年份" prop="student.enrollmentYear">
          <el-input v-model="formData.student.enrollmentYear" />
        </el-form-item>
        <el-form-item label="毕业年份" prop="student.generateYear">
          <el-input v-model="formData.student.generateYear" />
        </el-form-item>
      </template>

      <!-- 在职人员信息 -->
      <template v-else>
        <el-form-item label="公司" prop="employee.company">
          <el-input v-model="formData.employee.company" />
        </el-form-item>
        <el-form-item label="职位" prop="employee.position">
          <el-input v-model="formData.employee.position" />
        </el-form-item>
        <el-form-item label="行业" prop="employee.industry">
          <el-input v-model="formData.employee.industry" />
        </el-form-item>
        <el-form-item label="工作年限" prop="employee.workExperience">
          <el-input-number v-model="formData.employee.workExperience" :min="0" />
        </el-form-item>
      </template>

      <el-form-item label="手机号码" prop="phone">
        <el-input v-model="formData.phone" />
      </el-form-item>

      <el-form-item label="邮箱" prop="email">
        <el-input v-model="formData.email" />
      </el-form-item>

      <el-form-item label="参与目的" prop="purposes">
        <el-input v-model="formData.purposes" type="textarea" />
      </el-form-item>

      <el-form-item label="备注" prop="notes">
        <el-input v-model="formData.notes" type="textarea" />
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="$emit('update:visible', false)">取消</el-button>
      <el-button type="primary" @click="handleSubmit">提交</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">

import { defineProps, defineEmits, ref } from 'vue'
import type { FormInstance } from 'element-plus'

const formRef = ref<FormInstance>()
const formRules = {
  name: [
    {required: true, message: '姓名不能为空', trigger: 'blur'},
  ],
  phone: [
    {required: true, message: '电话号码不能为空', trigger: 'blur'},
    {pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码', trigger: 'blur'},
  ],
  email: [
    {required: true, message: '邮箱不能为空', trigger: 'blur'},
    {type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur'},
  ],
  purposes: [
    {required: true, message: '报名目的不能为空', trigger: 'blur'},
  ],

  'employee.company': [
    {required: true, message: '公司名称不能为空', trigger: 'blur'},
  ],
  'employee.position': [
    {required: true, message: '职位不能为空', trigger: 'blur'},
  ],
  'employee.industry': [
    {required: true, message: '行业不能为空', trigger: 'change'},
  ],
  'employee.workExperience': [
    {required: true, message: '工作经历不能为空', trigger: 'blur'},
  ],

  'student.school': [
    {required: true, message: '学校名称不能为空', trigger: 'blur'},
  ],
  'student.major': [
    {required: true, message: '专业不能为空', trigger: 'blur'},
  ],
  'student.enrollmentYear': [
    {required: true, message: '入学年份不能为空', trigger: 'blur'},
  ],
  'student.generateYear': [
    {required: true, message: '毕业年份不能为空', trigger: 'blur'},
  ],
};
const props =defineProps({
  visible: {
    type: Boolean,
    required: true
  },
  isEdit: {
    type: Boolean,
    default: false
  },
  formData: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['submit', 'update:visible'])

const handleSubmit = async () => {
  try {
    if (formRef.value) {
      const valid = await formRef.value.validate();
      if (valid) {
        // 使用 props.formData 来访问传递进来的表单数据
        emit('submit', props.formData);
      }
    }
  } catch (error) {
    console.error('表单验证失败:', error);
  }
};
</script>