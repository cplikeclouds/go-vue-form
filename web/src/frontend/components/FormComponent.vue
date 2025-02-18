<template>
  <el-form :model="formData" :rules="rules" ref="formRef" label-width="120px">
    <div class="form-container">
      <el-form-item label="姓名" prop="name">
        <el-input v-model="formData.name" placeholder="请输入姓名" clearable class="custom-input"></el-input>
      </el-form-item>
      <el-form-item label="电话号码" prop="phone">
        <el-input v-model="formData.phone" placeholder="请输入电话号码" clearable class="custom-input"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="formData.email" placeholder="请输入邮箱" clearable class="custom-input"></el-input>
      </el-form-item>
      <el-form-item label="报名目的" prop="purposes">
        <el-input v-model="formData.purposes" placeholder="请输入报名目的" clearable class="custom-input"></el-input>
      </el-form-item>
      <el-form-item label="其他需求">
        <el-input v-model="formData.notes" placeholder="其他需求..." clearable class="custom-input"></el-input>
      </el-form-item>


      <!-- 用户类型选择 -->
      <el-form-item label="角色类型" prop="userType">
        <el-radio-group v-model="formData.applicant_type">
          <el-radio value="student">学生</el-radio>
          <el-radio value="employee">在职人员</el-radio>
        </el-radio-group>
      </el-form-item>

      <!-- 动态显示字段 -->
      <template v-if="formData.applicant_type === 'employee'">
        <el-form-item label="公司名称" prop="employee.company">
          <el-input v-model="formData.employee.company" placeholder="请输入公司名称" clearable
                    class="custom-input"></el-input>
        </el-form-item>
        <el-form-item label="职位" prop="employee.position">
          <el-input v-model="formData.employee.position" placeholder="请输入职位" clearable
                    class="custom-input"></el-input>
        </el-form-item>
        <el-form-item label="行业" prop="employee.industry">
          <el-select v-model="formData.employee.industry" placeholder="选择行业" class="custom-input">
            <el-option label="IT/互联网" value="IT"></el-option>
            <el-option label="金融" value="finance"></el-option>
            <el-option label="教育" value="education"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作经历" prop="employee.workExperience">
          <el-input v-model="formData.employee.workExperience" placeholder="工作经历" clearable
                    class="custom-input"></el-input>
        </el-form-item>
      </template>

      <template v-if="formData.applicant_type === 'student'">
        <el-form-item label="学校名称" prop="student.school">
          <el-input v-model="formData.student.school" placeholder="请输入学校名称" clearable
                    class="custom-input"></el-input>
        </el-form-item>
        <el-form-item label="专业" prop="student.major">
          <el-input v-model="formData.student.major" placeholder="请输入专业" clearable class="custom-input"></el-input>
        </el-form-item>
        <el-form-item label="入学年份" prop="student.enrollmentYear">
          <el-date-picker
              v-model="formData.student.enrollmentYear"
              type="year"
              value-format="YYYY"
              placeholder="请选择入学年份"
              class="custom-input">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="毕业年份" prop="student.generateYear">
          <el-date-picker
              v-model="formData.student.generateYear"
              type="year"
              value-format="YYYY"
              placeholder="请选择毕业年份"
              class="custom-input">
          </el-date-picker>
        </el-form-item>

      </template>

      <el-form-item>
        <el-button @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </div>
  </el-form>
  <el-backtop :right="100" :bottom="100"/>
</template>

<script lang="ts" setup>
import {ref} from 'vue';

import api from '../../services/api.js'
import formModel from "../../common/models/formmodel.js";
import router from "../../router/index.js";
import {FormInstance, ElMessage} from 'element-plus'

const formData = formModel
const formRef = ref<FormInstance>();

const rules = {
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

const submitForm = async () => {
  const formEl = formRef.value;
  if (!formEl) return;
  try {
    await formEl.validate();
    const response = await api.post('/submit', formData.value);
    await router.push('/success');
    ElMessage({
      message: '提交成功!',
      type: 'success',
    })
    console.log('submit!', response);

  } catch (error) {


    ElMessage({
      message: '提交出错!',
      type: 'warning',
    })
    console.log('error submit!', error);
  }
};


const resetForm = () => {
  const formEl = formRef.value;
  if (!formEl) return;
  formEl.resetFields();
};

</script>

<style scoped>
/* 表单容器样式 */
.form-container {
  width: 80%;
  max-width: 600px; /* 可根据实际情况调整最大宽度 */
  margin: 0 auto; /* 水平居中 */
}

/* 定义自定义输入框类，设置宽度 */
.custom-input {
  width: 300px;
}

/* 调整 el-form-project-item 的上下间距 */
.el-form-item {
  margin-bottom: 20px;
}
</style>