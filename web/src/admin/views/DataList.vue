<template>
  <div class="data-list">
    <h2>数据列表</h2>
    <div class="operation-bar">
      <el-button type="primary" @click="handleCreate">新增报名者</el-button>
    </div>
    <el-table
        :data="tableData"
        style="width: 100%"
        :row-class-name="rowClassName"
        v-loading="loading"
    >
      <el-table-column prop="ID" label="ID" width="80">
        <template #default="scope">
          {{ renderValue(scope.row.ID) }}
        </template>
      </el-table-column>
      <el-table-column prop="applicant_type" label="类型" width="100">
        <template #default="scope">
          {{ renderValue(scope.row.applicant_type === 'student' ? '学生' : '在职人员') }}
        </template>
      </el-table-column>
      <el-table-column prop="name" label="姓名" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.name) }}
        </template>
      </el-table-column>
      <el-table-column prop="phone" label="手机号码" width="180">
        <template #default="scope">
          {{ renderValue(scope.row.phone) }}
        </template>
      </el-table-column>
      <el-table-column prop="email" label="邮箱" width="180">
        <template #default="scope">
          {{ renderValue(scope.row.email) }}
        </template>
      </el-table-column>
      <el-table-column prop="purposes" label="参与目的">
        <template #default="scope">
          {{ renderValue(scope.row.purposes) }}
        </template>
      </el-table-column>
      <el-table-column prop="notes" label="备注">
        <template #default="scope">
          {{ renderValue(scope.row.notes) }}
        </template>
      </el-table-column>
      <el-table-column label="学校" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.student?.school) }}
        </template>
      </el-table-column>
      <el-table-column label="专业" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.student?.major) }}
        </template>
      </el-table-column>
      <el-table-column label="入学年份" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.student?.enrollment_year) }}
        </template>
      </el-table-column>
      <el-table-column label="毕业年份" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.student?.generate_year) }}
        </template>
      </el-table-column>
      <el-table-column label="公司" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.employee?.company) }}
        </template>
      </el-table-column>
      <el-table-column label="职位" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.employee?.position) }}
        </template>
      </el-table-column>
      <el-table-column label="行业" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.employee?.industry) }}
        </template>
      </el-table-column>
      <el-table-column label="工作年限" width="120">
        <template #default="scope">
          {{ renderValue(scope.row.employee?.work_experience) }}
        </template>
      </el-table-column>
      <el-table-column prop="CreatedAt" label="创建时间" width="180">
        <template #default="scope">
          {{ renderValue(formatDate(scope.row.CreatedAt)) }}
        </template>
      </el-table-column>
      <el-table-column prop="UpdatedAt" label="更新时间" width="180">
        <template #default="scope">
          {{ renderValue(formatDate(scope.row.UpdatedAt)) }}
        </template>
      </el-table-column>
      <!-- 操作列 -->
      <el-table-column label="操作" width="150" fixed="right">
        <template #default="scope">
          <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
          <el-popconfirm
              title="确认删除该记录吗？"
              @confirm="handleDelete(scope.row.ID)"
          >
            <template #reference>
              <el-button size="small" type="danger">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <data-form-dialog
        v-model:visible="dialogVisible"
        :is-edit="isEditMode"
        :form-data="currentForm"
        @submit="handleSubmit"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { Applicant, createData, deleteData, fetchDataList, updateData } from '../api';
import DataFormDialog from '../components/DataFormDialog.vue';
import dayjs from 'dayjs'; // 用于日期格式化

const loading = ref(false);
const dialogVisible = ref(false);
const isEditMode = ref(false);
const tableData = ref<Applicant[]>([]);
const currentForm = ref<Applicant>({
  applicant_type: 'student',
  name: '',
  phone: '',
  email: '',
  purposes: '',
  notes: '',
  student: {
    school: '',
    major: '',
    enrollmentYear: ''
  },
  employee: {
    company: '',
    position: '',
    industry: '',
    workExperience: 0
  }
});

// 初始化加载数据
onMounted(async () => {
  await loadData();
});

// 加载数据
const loadData = async () => {
  try {
    loading.value = true;
    const response = await fetchDataList();
    console.log(response);
    tableData.value = response.data;
  } catch (error) {
    console.error('加载数据失败:', error);
  } finally {
    loading.value = false;
  }
};

// 格式化日期
const formatDate = (date: string) => {
  return dayjs(date).format('YYYY-MM-DD HH:mm:ss');
};

// 创建操作
const handleCreate = () => {
  currentForm.value = {
    applicant_type: 'student',
    name: '',
    phone: '',
    email: '',
    purposes: '',
    notes: '',
    student: {
      school: '',
      major: '',
      enrollmentYear: ''
    },
    employee: {
      company: '',
      position: '',
      industry: '',
      workExperience: 0
    }
  };
  isEditMode.value = false;
  dialogVisible.value = true;
};

// 编辑操作
const handleEdit = (row: Applicant) => {
  currentForm.value = JSON.parse(JSON.stringify(row));
  isEditMode.value = true;
  dialogVisible.value = true;
};

// 提交表单
const handleSubmit = async (formData: Applicant) => {
  try {
    if (isEditMode.value && formData.id) {
      await updateData(formData.id, formData);
    } else {
      await createData(formData);
    }
    await loadData();
    dialogVisible.value = false;
  } catch (error) {
    console.error('操作失败:', error);
  }
}

// 删除操作
const handleDelete = async (id: number) => {
  try {
    await deleteData(id);
    await loadData(); // 刷新数据
  } catch (error) {
    console.error('删除失败:', error);
  }
};

const renderValue = (value: any) => {
  if (value === null || value === undefined || value === '') {
    return '/';
  }
  return value;
};

// 行样式分类
const rowClassName = ({ row }: { row: Applicant }) => {
  return row.applicant_type === 'student' ? 'student-row' : 'employee-row';
};


// 导出方法供外部调用
const exposeMethods = {
  renderValue,
  handleCreate,
  handleEdit,
  handleDelete,
  handleSubmit
};
defineExpose(exposeMethods);
</script>

<style scoped>
/* 样式部分保持不变 */
</style>